package chain

// ServerManager stores all relevant data of our runtime
type ServerManager struct {
	Name        string
	BlockChains map[string]*Blockchain
}
