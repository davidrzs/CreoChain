// taken from https://github.com/Jeiwan/blockchain_go/blob/part_1/blockchain.go

package chain

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	Name        string
	accessToken string
	Blocks      []*Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain(name string) *Blockchain {
	return &Blockchain{name, "", []*Block{NewGenesisBlock()}}
}
