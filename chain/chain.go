// taken from https://github.com/Jeiwan/blockchain_go/blob/part_1/blockchain.go

package chain

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	name   string
	blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain(name string) *Blockchain {
	return &Blockchain{name, []*Block{NewGenesisBlock()}}
}

func SaveBlock() {

}
