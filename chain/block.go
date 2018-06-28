// Taken from https://github.com/Jeiwan/blockchain_go/blob/part_1/block.go

package chain

import (
	"bytes"
	"crypto/sha256"
	b64 "encoding/base64"
	"strconv"
	"time"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// setHash calculates and sets block hash
func (b *Block) setHash() {
	b.Hash = GetHash(b)
}

// Base64Hash gets the base64 encoding of the hash
func (b *Block) Base64Hash() {
	b64.StdEncoding.EncodeToString(b.Hash)
}

//GetHash returns the hash of a block.
func GetHash(b *Block) []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.setHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
