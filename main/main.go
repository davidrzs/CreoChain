package main

import (
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

// ServerManager stores all relevant data of our runtime
type ServerManager struct {
	name        string
	blockChains map[string]*chain.Blockchain
}

func initializePersistence() *bolt.DB {
	db := chain.InitializeDatabase(databaseName)

	err1 := chain.CreateBucket(db, metaInfoBucketName)
	err2 := chain.CreateBucket(db, chainBucketName)

	if err1 != nil || err2 != nil {
		panic("error ocurred while creating buckets")
	}
	return db
}

func main() {
	// begin database initialization
	db := initializePersistence()
	defer db.Close() //remember to close it at the end of program execution
	// end database initialization

	// begin variable assignment and reading in from database
	Data := &ServerManager{name: "main dataset", blockChains: make(map[string]*chain.Blockchain)}
	// end variable assignment and reading in from database

	// begin server
	server.Serve()
	//end server

	// begin debugging
	fmt.Println("Up and running")
	chain.Test()
	// end debugging

}
