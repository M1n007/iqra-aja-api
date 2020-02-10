package main

import (
	"fmt"
	server "iqra-aja-api/bin"
)

func main() {
	fmt.Printf("Server is running on port 9000")
	server.Init()
}
