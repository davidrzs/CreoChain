package server

import (
	"log"
	"net/http"
	"runtime"
	"strconv"

	"../chain"
	"../globalvariables"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

// Serve fires up the server.
func Serve(Data *globalvariables.ServerManager) {
	e := gin.Default()

	var m runtime.MemStats
	e.GET("/", func(c *gin.Context) {
		runtime.ReadMemStats(&m)
		stringToDisplay := WelcomeCreo + "\n" + "\n" + "Current Architecture: " + string(runtime.GOOS) + "\n" + "Current Memory Usage: " + strconv.Itoa(int(bToMb(m.TotalAlloc))) + "Mb" + "\n" + "Check out the documentation at https://davidrzs.github.io/CreoChain/"
		c.String(http.StatusOK, stringToDisplay)
	})

	// get the fullchain in JSON format of a non pretected chain
	e.GET("/v1/chain/:chainname/", func(c *gin.Context) {

		Data.Mutex.Lock()
		defer Data.Mutex.Unlock()

		bchain := chain.Chain{}

		if Data.Database.Where("name = ?", c.Param("chainname")).First(&bchain).RecordNotFound() {
			c.String(http.StatusNotFound, "Error 404. The chain you wanted to retrieve doesn't exist.")
		} else {
			blocks := []chain.Block{}
			Data.Database.Model(&bchain).Association("blocks").Find(&blocks)
			bchain.Blocks = blocks
			c.JSON(http.StatusOK, &bchain)
		}
	})

	// get a specific block in JSON format
	e.GET("/v1/chain/:chainname/block/:blockid", func(c *gin.Context) {
		Data.Mutex.Lock()
		defer Data.Mutex.Unlock()
		bblock := chain.Block{}
		bchain := chain.Chain{}
		if Data.Database.Where("name = ?", c.Param("chainname")).First(&bchain).RecordNotFound() {
			c.String(http.StatusNotFound, "Error 404. The chain you wanted to retrieve doesn't exist.")
		} else {
			if Data.Database.Where("chain_id = ? AND id_in_blockchain = ?", bchain.ChainID, c.Param("blockid")).First(&bblock).RecordNotFound() {
				c.String(http.StatusNotFound, "Error 404. The block you wanted to retrieve doesn't exist on the specified chain.")
			} else {
				c.JSON(http.StatusOK, &bblock)
			}
		}
	})
	/*
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
		/*
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
	if Data.Config.Server.Usessl == true {
		log.Fatal(autotls.Run(e, Data.Config.Server.Urls...))
	} else {
		e.Run(":8080")
	}
}
