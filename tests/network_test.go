package tests

import (
	"testing"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/network"
)

func TestPeerManager(t *testing.T) {
	pm := network.NewPeerManager()
	pm.AddPeer("127.0.0.1:8080")

	peers := pm.GetActivePeers()
	if len(peers) != 1 || peers[0] != "127.0.0.1:8080" {
		t.Errorf("Failed to add peer. Expected 1 peer, got %d", len(peers))
	}
}
