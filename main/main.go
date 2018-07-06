package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"../chain"
	"../globalvariables"
	"../persistence"
	"../server"
)

const (
	version = "V0.2"
)

func main() {
	//read config yaml
	content, err1 := ioutil.ReadFile("config.yml")
	if err1 != nil {
		fmt.Println("couldn't read config.yaml: \n Error: ")
		panic("couldn't read config.yaml: \n Error: \n" + err1.Error())
	}

	config, err2 := persistence.ParseYAML(string(content))
	if err2 != nil {
		panic("Couldn't parse contents of config.yaml: Error: " + err2.Error())
	}
	fmt.Println("this is ", config.Server.Globalauthcode)

	// begin variable assignment and reading in from database
	dbType, dbString := globalvariables.DatabaseConnectionString(config)
	db := chain.DbSetup(dbType, dbString)
	Data := &globalvariables.ServerManager{Mutex: &sync.Mutex{}, Name: "main dataset", Database: db, Config: config}
	// end variable assignment and reading in from database

	// begin debugging
	fmt.Println("Up and running")
	//chain.CreateNewBlockchain(Data, "othertest", "auth")
	//chain.AddBlockToChain(Data.Database, "chain1", "aT", "this is some sd for is")
	// end debugging

	// begin server
	server.Serve(Data) // running it in another thread
	//end server

	//beginning shutdown handling
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop
	fmt.Println("\nRecieved shutdown signal")
	db.Close() //remember to close it at the end of program execution
	fmt.Println("Closed Database")
	time.Sleep(500 * time.Millisecond)
	os.Exit(0)
	//ending shutdown handling

}
