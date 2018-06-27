package server

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"

	"../chain"
	"github.com/labstack/echo"
)

const welcomeCreo = "   _____                  ____  _            _        _           _          _____                          \r\n  / ____|                |  _ \\| |          | |      | |         (_)        / ____|                         \r\n | |     _ __ ___  ___   | |_) | | ___   ___| | _____| |__   __ _ _ _ __   | (___   ___ _ ____   _____ _ __ \r\n | |    | '__/ _ \\/ _ \\  |  _ <| |/ _ \\ / __| |/ / __| '_ \\ / _` | | '_ \\   \\___ \\ / _ \\ '__\\ \\ / / _ \\ '__|\r\n | |____| | |  __/ (_) | | |_) | | (_) | (__|   < (__| | | | (_| | | | | |  ____) |  __/ |   \\ V /  __/ |   \r\n  \\_____|_|  \\___|\\___/  |____/|_|\\___/ \\___|_|\\_\\___|_| |_|\\__,_|_|_| |_| |_____/ \\___|_|    \\_/ \\___|_|   \r\n                                                                                                            \r\n                           "

/*
# General Comments:

We have a versioned API -> we are currently v1
*/
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Serve fires up the server.
func Serve(data *chain.ServerManager) {
	e := echo.New()
	var m runtime.MemStats
	e.GET("/", func(c echo.Context) error {
		runtime.ReadMemStats(&m)
		stringToDisplay := welcomeCreo + "\n" + "\n" + "Current Architecture: " + string(runtime.GOOS) + "\n" + "Current Memory Usage: " + strconv.Itoa(int(bToMb(m.TotalAlloc))) + "Mb" + "\n" + "Check out the documentation at www.1234.com"
		return c.String(http.StatusOK, stringToDisplay)
	})

	// get the fullchain in JSON format of a non pretected chain
	e.GET("/v1/chain/:chainname/fullchain", func(c echo.Context) error {
		chain, isPresent := data.BlockChains[c.Param("chainname")]
		if !isPresent {
			return c.String(http.StatusNotFound, "Error 404. The chain you wanted to retrieve doesn't exist.")
		}
		var jsonString strings.Builder

		/*{"employees":[
		    { "firstName":"John", "lastName":"Doe" },
		    { "firstName":"Anna", "lastName":"Smith" },
		    { "firstName":"Peter", "lastName":"Jones" }
		]}*/
		buffer.WriteString("'blocks': [ \n")
		for _, block := range chain.Blocks {
			buffer.WriteString("{'prevBlockHash':'")
			buffer.WriteString(block.PrevBlockHash)
			buffer.WriteString("{'prevBlockHash':'")

			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
		buffer.WriteString("]")

		if err != nil {
			return c.String(http.StatusInternalServerError, "Error 500. JSON marshalling didn't work.")
		}
		return c.JSON(http.StatusOK, buffer.String())
	})

	/*
		// get a specific block in JSON format
		e.GET("/v1/chain/:chainid/sb/:blockid", singleBlock)

		// recalculate all hashes in a chain and verify if they match the ones stored
		e.GET("/v1/chain/:chainid/checkchainhashes", checkHashesChain)

		// check the hash of a single block
		e.POST("/v1/chain/:chainid/checkblockhash", checkBlockHash)

		// add a single block to the end of a blockchain
		e.POST("/v1/chain/:chainid/checkblockhash", addBlockToChain)

		e.POST("/v1/chain/:chainid/checkblockhash", addBlockToChain)
	*/
	e.Logger.Fatal(e.Start(":1235"))
}
