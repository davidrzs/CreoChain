package main

import (
	"fmt"
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
	version            = "v0.1"
	metaInfoBucketName = "metainfobucket"
	chainBucketName    = "chainbucket"
	databaseName       = "creoDB.db"

	yamlString = `
database: keyvalue
server:
  globalAuthcodes: ["gs123","sadjksad"]
  usessl: true
`
)

func main() {
	// begin variable assignment and reading in from database
	db := chain.DbSetup("mysql", "david:password@/godb?charset=utf8&parseTime=True&loc=Local")
	Data := &globalvariables.ServerManager{Mutex: &sync.Mutex{}, Name: "main dataset", Database: db}
	Config, err := persistence.ParseYAML(yamlString)
	if err != nil {
		fmt.Println("An error occurred reading the Yaml file. Please fix this")
		fmt.Println(err)
		os.Exit(1)
	}
	// end variable assignment and reading in from database

	// begin debugging
	fmt.Println("Up and running")
	// end debugging

	// begin server
	go server.Serve(Data, Config) // running it in another thread
	//end server

	//beginning shutdown handling
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop
	fmt.Println("Recieved shutdown signal")
	close(quit) // closing the database saving channel stopping database access.
	db.Close()  //remember to close it at the end of program execution
	fmt.Println("Closed Database")
	time.Sleep(2 * time.Second)
	os.Exit(0)
	//ending shutdown handling

}
