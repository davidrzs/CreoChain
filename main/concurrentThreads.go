package main

import (
	"../chain"
	bolt "github.com/coreos/bbolt"
)

//PeriodicallySaveToDatabase does exactly what you think and takes a delay specified in seconds.
func PeriodicallySaveToDatabase(delay int, data *chain.ServerManager, database *bolt.DB) bool {
	//  case <-stop: // triggered when the stop channel is closed
	//break loop // exit

}
