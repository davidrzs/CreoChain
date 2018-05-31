package main

import (
	"../chain"
	"../server"
	"fmt"
)

func main() {
	server.Serve()
	fmt.Println("Up and running")

}
