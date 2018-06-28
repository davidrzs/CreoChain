package server

import (
	b64 "encoding/base64"
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"../chain"
	"github.com/labstack/echo"
)

const welcomeCreo = "   _____                  ____  _            _        _           _          _____                          \r\n  / ____|                |  _ \\| |          | |      | |         (_)        / ____|                         \r\n | |     _ __ ___  ___   | |_) | | ___   ___| | _____| |__   __ _ _ _ __   | (___   ___ _ ____   _____ _ __ \r\n | |    | '__/ _ \\/ _ \\  |  _ <| |/ _ \\ / __| |/ / __| '_ \\ / _` | | '_ \\   \\___ \\ / _ \\ '__\\ \\ / / _ \\ '__|\r\n | |____| | |  __/ (_) | | |_) | | (_) | (__|   < (__| | | | (_| | | | | |  ____) |  __/ |   \\ V /  __/ |   \r\n  \\_____|_|  \\___|\\___/  |____/|_|\\___/ \\___|_|\\_\\___|_| |_|\\__,_|_|_| |_| |_____/ \\___|_|    \\_/ \\___|_|   \r\n                                                                                                            \r\n                           "

// SingleHashCheck is the result of a single has check.
type SingleHashCheck struct {
	hash1 string
	hash2 string
	same  bool
}

//HashResult is the struct we return as json when we check all hashes.
type HashResult struct {
	HashesOk      bool
	Details       *[]SingleHashCheck
	DiscrepancyID int
}

// AreByteArraysEqual checks if two byte arrays are equal.
func AreByteArraysEqual(a []byte, b []byte) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

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

		return c.JSON(http.StatusOK, &chain)
	})

	// get a specific block in JSON format
	e.GET("/v1/chain/:chainname/sb/:blockid", func(c echo.Context) error {
		chain, isPresent := data.BlockChains[c.Param("chainname")]
		blockid, err := strconv.Atoi(c.Param("blockid"))
		if !isPresent {
			return c.String(http.StatusNotFound, "Error 404. The chain you wanted to retrieve doesn't exist.")
		}

		if err != nil {
			return c.String(http.StatusInternalServerError, "Error 500. Couln't parse block id.")
		}
		fmt.Print(len(chain.Blocks))
		if len(chain.Blocks)-1 < blockid {
			return c.String(http.StatusInternalServerError, "Error 500. The block id index exceeds the number of elements in the blockchain")
		}

		return c.JSON(http.StatusOK, &chain.Blocks[blockid])
	})

	// recalculate all hashes in a chain and verify if they match the ones stored
	e.GET("/v1/chain/:chainname/checkchainhashes", func(c echo.Context) error {
		bchain, isPresent := data.BlockChains[c.Param("chainname")]
		if !isPresent {
			return c.String(http.StatusNotFound, "Error 404. The chain you wanted to retrieve doesn't exist.")
		}
		discrepancy := false
		discrepancyid := 0
		var listOfHashChecks []SingleHashCheck
		for idx, block := range bchain.Blocks {
			fmt.Println(block)
			origHash := block.Hash
			newHash := chain.GetHash(block)
			currentDiscrepancy := !AreByteArraysEqual(newHash, origHash)

			ck := SingleHashCheck{hash1: string(origHash[:]), hash2: b64.StdEncoding.EncodeToString(newHash[:]), same: currentDiscrepancy}
			fmt.Println(ck)
			listOfHashChecks = append(listOfHashChecks, ck)
			if currentDiscrepancy == true {
				discrepancyid = idx
				discrepancy = true
			}
		}

		return c.JSON(http.StatusOK, &HashResult{HashesOk: !discrepancy,
			Details: &listOfHashChecks, DiscrepancyID: discrepancyid})
	})
	/*
		// check the hash of a single block
		e.POST("/v1/chain/:chainid/checkblockhash", checkBlockHash)

		// add a single block to the end of a blockchain
		e.POST("/v1/chain/:chainid/checkblockhash", addBlockToChain)

	*/
	e.Logger.Fatal(e.Start(":1235"))
}
