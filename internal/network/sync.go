package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/blockchain"
)

type Synchronizer struct {
	Blockchain *blockchain.Blockchain
	PeerManager *PeerManager
}

func NewSynchronizer(bc *blockchain.Blockchain, pm *PeerManager) *Synchronizer {
	return &Synchronizer{
		Blockchain: bc,
		PeerManager: pm,
	}
}

func (sync *Synchronizer) SyncBlockchain() {
	activePeers := sync.PeerManager.GetActivePeers()

	for _, peer := range activePeers {
		go func(peer string) {
			resp, err := http.Get(fmt.Sprintf("http://%s/blocks", peer))
			if err != nil {
				fmt.Printf("Failed to fetch blockchain from peer %s: %v\n", peer, err)
				return
			}
			defer resp.Body.Close()

			var receivedBlocks []*blockchain.Block
			if err := json.NewDecoder(resp.Body).Decode(&receivedBlocks); err != nil {
				fmt.Printf("Failed to decode blockchain from peer %s: %v\n", peer, err)
				return
			}

			if len(receivedBlocks) > len(sync.Blockchain.Blocks) {
				sync.Blockchain.Blocks = receivedBlocks
				fmt.Printf("Blockchain synchronized with peer %s\n", peer)
			}
		}(peer)
	}
}

func (sync *Synchronizer) BroadcastBlock(block *blockchain.Block) {
	activePeers := sync.PeerManager.GetActivePeers()

	for _, peer := range activePeers {
		go func(peer string) {
			msg := Message{
				Type:    "NEW_BLOCK",
				Payload: block,
			}
			if err := SendMessage(peer, msg); err != nil {
				fmt.Printf("Failed to broadcast block to peer %s: %v\n", peer, err)
			}
		}(peer)
	}
}
