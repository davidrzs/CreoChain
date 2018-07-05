// // taken from https://github.com/Jeiwan/blockchain_go/blob/part_1/blockchain.go

// package chain

// import (
// 	"fmt"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// )

// // Blockchain keeps a sequence of Blocks
// type Blockchain struct {
// 	gorm.Model
// 	Name        string `gorm:"unique"`
// 	AccessToken string
// 	Blocks      []*Block
// }

// //Product ...
// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

// // AddBlock saves provided data as a block in the blockchain
// func (bc *Blockchain) AddBlock(data string) {
// 	prevBlock := bc.Blocks[len(bc.Blocks)-1]
// 	newBlock := NewBlock(data, prevBlock.Hash)
// 	bc.Blocks = append(bc.Blocks, newBlock)
// }

// // NewBlockchain creates a new Blockchain with genesis Block
// func NewBlockchain(name string, authcode string) *Blockchain {
// 	return &Blockchain{Name: name, AccessToken: authcode, Blocks: []*Block{NewGenesisBlock()}}
// }

// //Tt ..
// func Tt() {

// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()
// 	db.DropTableIfExists(&User{},
// 	// Migrate the schema
// 	db.AutoMigrate(&Product{})
// 	db.AutoMigrate(&Block{})
// 	db.AutoMigrate(&Blockchain{})

// 	// Create
// 	db.Create(&Blockchain{Name: "test", AccessToken: "authcode", Blocks: []*Block{NewGenesisBlock()}})
// 	// err2 := db.Create(&Blockchain{Name: "test", AccessToken: "sadasd", Blocks: []Block{NewGenesisBlock()}})

// 	// Read
// 	var product Product
// 	db.First(&product, 1)                  // find product with id 1
// 	db.Last(&product, "code = ?", "L1212") // find product with code l1212
// 	var chain Blockchain
// 	db.Where("name = ?", "test").First(&chain)
// 	fmt.Println(chain)

// 	chain.AddBlock("dsdfsdf")
// 	db.Model(&chain).Update()
// 	fmt.Println(chain)

// }
