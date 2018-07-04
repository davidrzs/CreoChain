package server

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"../chain"
	"../persistence"
	"github.com/gin-gonic/gin"
)

// Serve fires up the server.
func Serve(data *chain.ServerManager, config *persistence.YAMLReader) {
	e := gin.Default()

	var m runtime.MemStats
	e.GET("/", func(c *gin.Context) {
		runtime.ReadMemStats(&m)
		stringToDisplay := WelcomeCreo + "\n" + "\n" + "Current Architecture: " + string(runtime.GOOS) + "\n" + "Current Memory Usage: " + strconv.Itoa(int(bToMb(m.TotalAlloc))) + "Mb" + "\n" + "Check out the documentation at www.1234.com"
		c.String(http.StatusOK, stringToDisplay)
	})

	// get the fullchain in JSON format of a non pretected chain
	e.GET("/v1/chain/:chainname/", func(c *gin.Context) {
		data.Mutex.Lock()
		chain, isPresent := data.BlockChains[c.Param("chainname")]
		defer data.Mutex.Unlock()

		if !isPresent {
			c.String(http.StatusNotFound, "Error 404. The chain you wanted to retrieve doesn't exist.")
		}

		c.JSON(http.StatusOK, &chain)

	})

	// get a specific block in JSON format
	e.GET("/v1/chain/:chainname/block/:blockid", func(c *gin.Context) {
		data.Mutex.Lock()
		chain, isPresent := data.BlockChains[c.Param("chainname")]
		defer data.Mutex.Unlock()

		blockid, err := strconv.Atoi(c.Param("blockid"))
		if !isPresent {
			c.String(http.StatusNotFound, "Error 404. The chain you wanted to retrieve doesn't exist.")
		}

		if err != nil {
			c.String(http.StatusInternalServerError, "Error 500. Couln't parse block id.")
		}
		fmt.Print(len(chain.Blocks))
		if len(chain.Blocks)-1 < blockid {
			c.String(http.StatusInternalServerError, "Error 500. The block id index exceeds the number of elements in the blockchain")
		}

		c.JSON(http.StatusOK, &chain.Blocks[blockid])
	})

	// recalculate all hashes in a chain and verify if they match the ones stored
	e.GET("/v1/chain/:chainname/checkchainhashes", func(c *gin.Context) {
		data.Mutex.Lock()
		defer data.Mutex.Unlock()

		bchain, isPresent := data.BlockChains[c.Param("chainname")]
		if !isPresent {
			c.String(http.StatusNotFound, "Error 404. The chain you wanted to retrieve doesn't exist.")
		}
		discrepancy := false
		discrepancyid := 0
		for idx, block := range bchain.Blocks {
			fmt.Println(block)
			origHash := block.Hash
			newHash := chain.GetHash(block)
			currentDiscrepancy := newHash == origHash
			if currentDiscrepancy == true {
				discrepancyid = idx
				discrepancy = true
			}
		}
		ret := &HashResult{HashesOk: !discrepancy, DiscrepancyID: discrepancyid}
		c.JSON(http.StatusOK, ret)
	})

	// add a single block to the end of a blockchain
	e.POST("/v1/chain/:chainname/new", func(c *gin.Context) {
		data.Mutex.Lock()
		defer data.Mutex.Unlock()

		bchain, isbPresent := data.BlockChains[c.Param("chainname")]
		if !isbPresent {
			c.String(http.StatusNotFound, "Error 404. The chain you wanted to retrieve doesn't exist.")
		}
		u := new(BlockAdder)
		fmt.Println(u.Content)
		if err := c.Bind(u); err != nil {
			c.String(http.StatusInternalServerError, "Error 500. Something is wrong with the JSON you supplied. \n Couldn't parse it correctly. \n Please consult the documentation or report a bug.")
		}
		if u.Authentication == bchain.AccessToken {
			bchain.AddBlock(u.Content)
			c.String(http.StatusOK, "Success! Block Added")
		}

		c.String(http.StatusUnauthorized, "Your authentication token was wrong. No write permission granted. The block could not be addded")
	})

	/*
		// check the hash of a single block
		e.POST("/v1/chain/:chainid/checkblockhash", checkBlockHash)
	*/
	e.Run(":8080")
}
