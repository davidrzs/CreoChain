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

	// add a single block to the end of a blockchain
	e.POST("/v1/chain/:chainname/", func(c *gin.Context) {
		Data.Mutex.Lock()
		defer Data.Mutex.Unlock() //this function can definitely be cleaned up a bit more
		bchain := chain.Chain{}
		if Data.Database.Where("name = ?", c.Param("chainname")).First(&bchain).RecordNotFound() {
			c.String(http.StatusNotFound, "Error 404. The chain you wanted to add a block to doesn't exist.")
		} else {
			//reading in data
			ab := AddBlock{}
			err1 := c.Bind(&ab)
			if err1 == nil {
				err2, errString := chain.AddBlockToChain(Data.Database, bchain.Name, ab.Authcode, ab.Data)
				if err2 == false {
					//so we were able to add it to the blockchain:
					c.String(http.StatusOK, "Block Added")
				} else {
					// so authorization probably failed
					c.String(http.StatusUnauthorized, "Your authentication token was wrong. No write permission granted. The block could not be addded \n If this is not the problem, there might be something wrong with the database. Check the logs"+errString)
				}
			} else {
				c.String(http.StatusInternalServerError, err1.Error())
			}
		}

	})

	e.POST("/v1/chain/", func(c *gin.Context) {
		Data.Mutex.Lock()
		defer Data.Mutex.Unlock() //this function can definitely be cleaned up a bit more

		//reading in data
		ac := AddChain{}
		err1 := c.Bind(&ac)

		if err1 == nil {
			accessCorrect, err2 := chain.CreateNewBlockchain(Data, ac.Name, ac.Globalauthcode)
			if accessCorrect == true {
				c.String(http.StatusOK, "Block Added")
			} else {
				c.String(http.StatusUnauthorized, "Your authentication token was wrong. No write permission granted. The block could not be addded \n If this is not the problem, there might be something wrong with the database. Check the logs"+err2)
			}
		} else {
			c.String(http.StatusInternalServerError, err1.Error())
		}

	})

	if Data.Config.Server.Usessl == true {
		log.Fatal(autotls.Run(e, Data.Config.Server.Urls...))
	} else {
		e.Run(":8080")
	}
}
