package network

import (
	"fmt"
	"sync"
)

type Peer struct {
	Address string
	IsActive bool
}

type PeerManager struct {
	Peers map[string]*Peer
	mutex sync.Mutex
}

func NewPeerManager() *PeerManager {
	return &PeerManager{
		Peers: make(map[string]*Peer),
	}
}

func (pm *PeerManager) AddPeer(address string) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if _, exists := pm.Peers[address]; !exists {
		pm.Peers[address] = &Peer{
			Address:  address,
			IsActive: true,
		}
		fmt.Printf("Peer added: %s\n", address)
	}
}

func (pm *PeerManager) RemovePeer(address string) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	delete(pm.Peers, address)
	fmt.Printf("Peer removed: %s\n", address)
}

func (pm *PeerManager) GetActivePeers() []string {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	activePeers := []string{}
	for address, peer := range pm.Peers {
		if peer.IsActive {
			activePeers = append(activePeers, address)
		}
	}
	return activePeers
}
