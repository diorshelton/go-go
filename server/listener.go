package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
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

	connChan := make(chan net.Conn, workers)

	var wg sync.WaitGroup
	startWorkers(&wg, connChan, workers)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error:%v", err)
			break
		}

		connChan <- conn
	}
	//Close before wait to prevent deadlock between channel closing and workers finishing tasks
	close(connChan)
	wg.Wait()

}

func startWorkers(wg *sync.WaitGroup, channel chan net.Conn, workers int) {
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for conn := range channel {
				handleConn(conn)
			}
		}()
	}
}

func handleConn(conn net.Conn) {
}
