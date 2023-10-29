package main

import (
	"kaurapuuro/net"
	"log"
)

func main() {

	log.Println("Starting server...")
	net.StartServer(":8080")
}
