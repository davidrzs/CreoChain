package server

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

/*
# General Comments:

We have a versioned API -> we are currently v1
*/

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// get the fullchain in JSON format
	e.GET("/v1/chain/:chainid/fullchain", fullChain)

	// get a specific block in JSON format
	e.GET("/v1/chain/:chainid/sb/:blockid", singleBlock)

	// recalculate all hashes in a chain and verify if they match the ones stored
	e.GET("/v1/chain/:chainid/checkchainhashes", checkHashesChain)

	// check the hash of a single block
	e.POST("/v1/chain/:chainid/checkblockhash", checkBlockHash)

	// add a single block to the end of a blockchain
	e.POST("/v1/chain/:chainid/checkblockhash", addBlockToChain)

	e.POST("/v1/chain/:chainid/checkblockhash", addBlockToChain)

	e.Logger.Fatal(e.Start(":1235"))
}

func fullChain(c echo.Context) error {
	// User ID from path `users/:id`
	id, err1 := strconv.Atoi(c.Param("chainid"))
	if err1 != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	u := "hello"
	return c.JSON(http.StatusOK, u)
}
