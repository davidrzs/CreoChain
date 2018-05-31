package main

import (
	"../chain"
	//"../server"
	bolt "github.com/coreos/bbolt"

	"fmt"
)

const (
	version            = "v0.1"
	metaInfoBucketName = "metainfobucket"
	chainBucketName    = "chainbucket"
	databaseName       = "creoDB"
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

func main() {
	db := initializePersistence()
	defer db.Close() //remember to close it at the end of program execution
	//server.Serve()
	fmt.Println("Up and running")

}
