package main

import (
	"flag"
	"go-go/cache"
	"go-go/handlers"
	"go-go/server"
)

func main() {

	portPtr := flag.String("port", "8080", "port address for server connection")
	workersPtr := flag.Int("workers", 4, "port number")

	flag.Parse()

	var router server.Router

	store := cache.New()

	router.Register("GET", "/keys", handlers.HandleGetAll(store))
	router.Register("GET", "/keys/:id", handlers.HandleGet(store))
	router.Register("PUT", "/keys/:id", handlers.HandlePut(store))
	router.Register("DELETE", "/keys/:id", handlers.HandleDelete(store))

	server.Run(*portPtr, *workersPtr, &router)
}
