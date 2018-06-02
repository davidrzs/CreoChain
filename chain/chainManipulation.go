package chain

import "fmt"

//Test ...

// CheckIfChainExists checks if a chain with a given name can be found in our database
func CheckIfChainExists(name string) bool {
	return false
}

func GetWholeChain(name string) *Blockchain {
	return nil
}

func GetBlockOfChain() *Block {
	return nil
}

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
