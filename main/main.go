package main

import (
	"../server"
	"fmt"
)

func main() {
	server.Serve()
	fmt.Println("Up and running")

}
