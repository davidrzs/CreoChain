// // Taken from https://github.com/Jeiwan/blockchain_go/blob/part_1/block.go

// package chain

// import (
// 	"bytes"
// 	"crypto/sha256"
// 	b64 "encoding/base64"
// 	"strconv"
// 	"time"

// 	"github.com/jinzhu/gorm"
// )
//
// // Block keeps block headers
// type Block struct {
// 	gorm.Model
// 	Timestamp     int64
// 	Data          string
// 	PrevBlockHash string
// 	Hash          string
// 	BlockchainID  uint
// }

// // setHash calculates and sets block hash
// func (b *Block) setHash() {
// 	b.Hash = GetHash(b)
// }

// //GetHash returns the hash of a block.
// func GetHash(b *Block) string {
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
// 	headers := bytes.Join([][]byte{[]byte(b.PrevBlockHash), []byte(b.Data), timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)
// 	return b64.StdEncoding.EncodeToString(hash[:])
// }

// // NewBlock creates and returns Block
// func NewBlock(data string, prevBlockHash string) *Block {
// 	block := &Block{Timestamp: time.Now().Unix(), Data: data, PrevBlockHash: prevBlockHash, Hash: ""}
// 	block.setHash()
// 	return block
// }

// // NewGenesisBlock creates and returns genesis Block
// func NewGenesisBlock() *Block {
// 	return NewBlock("Genesis Block", "")
// }
