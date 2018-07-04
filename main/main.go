package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"fmt"

	"../chain"
	"../persistence"
	"../server"
	bolt "github.com/coreos/bbolt"
)

const (
	version            = "v0.1"
	metaInfoBucketName = "metainfobucket"
	chainBucketName    = "chainbucket"
	databaseName       = "creoDB.db"

	yamlString = `
database: keyvalue
server:
  authcodes: ["gs123","sadjksad"]
  usessl: true
`
)

func initializePersistence() *bolt.DB {
	db := chain.InitializeDatabase(databaseName)

	err1 := chain.CreateBucket(db, metaInfoBucketName)
	err2 := chain.CreateBucket(db, chainBucketName)

	if err1 != nil || err2 != nil {
		panic("error occurred while creating buckets")
	}
	return db
}

func main() {
	// begin database initialization
	db := initializePersistence()
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// handling saving to database
			case <-quit:
				fmt.Println("Gracefully stopped thread saving to database")
				ticker.Stop()
				return
			}
		}
	}()
	// end database initialization

	// begin variable assignment and reading in from database
	Data := &chain.ServerManager{Mutex: &sync.Mutex{}, Name: "main dataset", BlockChains: make(map[string]*chain.Blockchain)}
	Config, err := persistence.ParseYAML(yamlString)
	if err != nil {
		fmt.Println("An error occurred reading the Yaml file. Please fix this")
		fmt.Println(err)
		os.Exit(1)
	}
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
	go server.Serve(Data, Config) // running it in another thread
	//end server

	//beginning shutdown handling
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop
	close(quit) // closing the database saving channel stopping database access.
	db.Close()  //remember to close it at the end of program execution
	fmt.Println("Closing all programs")
	time.Sleep(2 * time.Second)
	os.Exit(0)
	//ending shutdown handling

}
