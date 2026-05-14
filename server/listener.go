package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(port string, workers int) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Panicf("Server error: %v", err)
	}
	defer listener.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigChan

		fmt.Printf("\nShutting down server %v", sig)

		listener.Close()
	}()

	pool := NewPool(workers, &Router{})

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error:%v", err)
			break
		}

		pool.Submit(conn)
	}
	pool.Shutdown()
}
