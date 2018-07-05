package chain

/*
// GetBlockOfChainByID gets a specific block of a specific chain
func GetBlockOfChainByID(chain *Blockchain, sindex int) Block {
	for index, block := range chain.Blocks {
		if index == sindex {
			return block
		}
	}
	return nil // since we didnt find any corresponding block
}

// GetBlockOfChainByContent returns a block containing a certain string in its data
func GetBlockOfChainByContent(chain *Blockchain, hash string) Block {
	for _, block := range chain.Blocks { // this method should be checkd for correctness
		hashstring := string(block.Hash[:])
		if hashstring == hash {
			return block
		}
	}
	return nil // since we didnt find any corresponding block
}

// GetBlockOfChainByHash returns the block corresponding to a hash
func GetBlockOfChainByHash(chain *Blockchain, search string) Block {
	for _, block := range chain.Blocks {
		datastring := string(block.Data[:])
		if strings.Contains(datastring, search) {
			return block
		}
	}
	return nil // since we didnt find any corresponding block
}

//StoreBlockChain ..
func StoreBlockChain(chain *Blockchain, name string) error {
	return nil
}
*/
// //Test ...
// func Test() {
// 	bc := NewBlockchain("test", "")

// 	bc.AddBlock("Send 1 BTC to Ivan")
// 	bc.AddBlock("Send 2 more BTC to Ivan")

// 	for _, block := range bc.Blocks {
// 		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
// 		fmt.Printf("Data: %s\n", block.Data)
// 		fmt.Printf("Hash: %x\n", string(block.Hash))
// 		fmt.Println()
// 	}
// }
