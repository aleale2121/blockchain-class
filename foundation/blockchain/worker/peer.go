package worker

import (
	"github.com/ardanlabs/blockchain/foundation/blockchain/peer"
)



// addNewPeers takes the list of known peers and makes sure they are included
// in the nodes list of know peers.
func (w *Worker) addNewPeers(knownPeers []peer.Peer) error {
	w.evHandler("worker: runPeerUpdatesOperation: addNewPeers: started")
	defer w.evHandler("worker: runPeerUpdatesOperation: addNewPeers: completed")

	for _, peer := range knownPeers {

		// Don't add this running node to the known peer list.
		if peer.Match(w.state.Host()) {
			continue
		}

		// Only log when the peer is new.
		if w.state.AddKnownPeer(peer) {
			w.evHandler("worker: runPeerUpdatesOperation: addNewPeers: add peer nodes: adding peer-node %s", peer.Host)
		}
	}

	return nil
}
