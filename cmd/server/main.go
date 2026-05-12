package main

import (
	"flag"
	"go-go/server"
)

func main() {

	portPtr := flag.String("port", "8080", "port address for server connection")
	workersPtr := flag.Int("workers", 4, "port number")

	flag.Parse()

	server.Run(*portPtr, *workersPtr)
}
