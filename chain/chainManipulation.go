package chain

import "fmt"

// CheckIfChainExists checks if a chain with a given name can be found in our database
func CheckIfChainExists(name string) bool {
	return false
}

// GetWholeChain returns a pointer to a blockchain loaded from the storage.
func GetWholeChain(name string) *Blockchain {
	return nil
}

// GetBlockOfChain gets a specific block of a specific chain
func GetBlockOfChain() *Block {
	return nil
}

//StoreBlockChain ..
func StoreBlockChain(chain *Blockchain, name string) error {
	return nil
}

//AddBlockToChain ..
func AddBlockToChain() error {
	return nil
}

// DeleteBlockChain deletes an entire blockchain from the database. Use with caution.
func DeleteBlockChain(name string) error {
	return nil
}

// DeleteBlockFromChain should NEVER be used since it would mess up the hashes, only useful for debugging.
func DeleteBlockFromChain() error {
	return nil
}

//Test ...
func Test() {
	bc := NewBlockchain("test")

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
