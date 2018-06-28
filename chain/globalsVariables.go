package chain

import "sync"

// ServerManager stores all relevant data of our runtime
type ServerManager struct {
	Mutex       *sync.Mutex
	Name        string
	BlockChains map[string]*Blockchain
}
