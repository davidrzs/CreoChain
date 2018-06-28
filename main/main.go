package main

import (
	"sync"

	"../chain"
	"../server"
	bolt "github.com/coreos/bbolt"

	"fmt"
)

const (
	version            = "v0.1"
	metaInfoBucketName = "metainfobucket"
	chainBucketName    = "chainbucket"
	databaseName       = "creoDB.db"
)

func initializePersistence() *bolt.DB {
	db := chain.InitializeDatabase(databaseName)

	err1 := chain.CreateBucket(db, metaInfoBucketName)
	err2 := chain.CreateBucket(db, chainBucketName)

	if err1 != nil || err2 != nil {
		panic("error ocurred while creating buckets")
	}
	return db
}

func cleanup() {
	close(number)
}

func main() {
	// begin database initialization
	db := initializePersistence()
	defer db.Close() //remember to close it at the end of program execution
	// end database initialization

	// begin variable assignment and reading in from database
	Data := &chain.ServerManager{Mutex: &sync.Mutex{}, Name: "main dataset", BlockChains: make(map[string]*chain.Blockchain)}
	// end variable assignment and reading in from database

	// begin debugging
	fmt.Println("Up and running")
	chain.Test()
	Data.BlockChains["test"] = chain.NewBlockchain("test", "autheR3")
	Data.BlockChains["test"].AddBlock("Send 2 more BTC to Ivan")

	Data.BlockChains["test"].AddBlock("Hello, this is me")
	Data.BlockChains["test"].AddBlock("another one")
	// end debugging

	// begin server
	server.Serve(Data)
	//end server

}
