package server

import (
	"net"
	"sync"
)

type Pool struct {
	connCh chan net.Conn
	wg     sync.WaitGroup
	router *Router
}

func NewPool(workers int, router *Router) *Pool {
	// TODO
	return nil
}

func (p *Pool) Submit(conn net.Conn) {
	// TODO
}

func (p *Pool) Shutdown() {
	// TODO
}
