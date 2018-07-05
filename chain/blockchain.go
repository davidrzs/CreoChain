package chain

import (
	"bytes"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Chain struct {
	ChainID     int    `gorm:"primary_key"`
	Name        string `gorm:"unique"`
	AccessToken string
	Blocks      []Block `gorm:"ForeignKey:ChainId"`
}

type Block struct {
	BlockID       int `gorm:"primary_key"`
	Timestamp     int64
	Data          string
	PrevBlockHash string
	Hash          string
	ChainId       int
}

type Config struct {
	ConfigID int `gorm:"primary_key"`
}

// setHash calculates and sets block hash
func (b *Block) setHash() {
	b.Hash = GetHash(b)
}

//GetHash returns the hash of a block.
func GetHash(b *Block) string {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{[]byte(b.PrevBlockHash), []byte(b.Data), timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return b64.StdEncoding.EncodeToString(hash[:])
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{Timestamp: time.Now().Unix(), Data: data, PrevBlockHash: prevBlockHash, Hash: ""}
	block.setHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}

func runChainTest() {

	db := DbSetup("mysql", "david:password@/godb?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	//db.DropTableIfExists(&Block{})
	//db.DropTableIfExists(&Chain{})
	/*Blockchain1 := Chain{Name: "chain1", AccessToken: "aT", Blocks: []Block{
		{Timestamp: time.Now().Unix(), Data: "block1", PrevBlockHash: "", Hash: ""},
		{Timestamp: time.Now().Unix(), Data: "block2", PrevBlockHash: "", Hash: ""}}}

	db.Create(&Blockchain1)
	*/
	bchain := &Chain{}

	fmt.Println(AddBlockToChain(db, "chain1", "aT", "this is some data"))

	fmt.Println("bchain", bchain.ChainID)

}

//AddBlockToChain adds a block to a blockchain by saving it correctly in the database.
func AddBlockToChain(db *gorm.DB, chainName string, authCode string, data string) (bool, string) {

	err := false
	errString := ""

	//find the chain we are looking for.
	bchain := &Chain{}

	if db.Where("name=?", chainName).First(&bchain).RecordNotFound() {
		err = true
		errString = "No chain with the name " + chainName + " found. "
		return err, errString
	}

	//check if we are authorized to add a block
	var trimmedAccessToken string
	if runtime.GOOS == "windows" {
		trimmedAccessToken = strings.TrimRight(authCode, "\r\n")
	} else {
		trimmedAccessToken = strings.TrimRight(authCode, "\n")
	}

	if bchain.AccessToken != trimmedAccessToken {
		err = true
		errString = "AccessToken for chain " + chainName + " isn't correct. "
		return err, errString
	}

	//creating the block
	bchainBlocks := &[]Block{}
	db.Model(&bchain).Association("Blocks").Find(bchainBlocks)
	fmt.Println(bchainBlocks)
	prevBlockHash := (*bchainBlocks)[len(*bchainBlocks)-1].Hash
	blockToAdd := NewBlock(data, prevBlockHash)
	blockToAdd.setHash()

	//add it to the database
	db.Model(&bchain).Association("Blocks").Append(blockToAdd)
	return false, ""
}

// DbSetup establishes a databse connection and returns it.
func DbSetup(dbType string, connectionString string) *gorm.DB {
	db, err := gorm.Open(dbType, connectionString)
	if err != nil {
		fmt.Println("Couldn't connect to the database with the credentials you supplied.")
		panic(err.Error())
	}
	db.AutoMigrate(&Chain{}, &Block{}, &Config{})                                          // make sure we've got the schema loaded
	db.Model(&Block{}).AddForeignKey("chain_id", "chains(chain_id)", "CASCADE", "CASCADE") // add foregin key
	return db
}