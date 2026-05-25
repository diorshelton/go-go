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

	connCh := make(chan net.Conn, workers)

	pool := &Pool{connCh: connCh, router: router}

	for i := 1; i <= workers; i++ {
		pool.wg.Add(1)
		go func() {
			defer pool.wg.Done()
			for conn := range connCh {
				pool.handleConn(conn)
			}
		}()
	}

	return pool
}

func (p *Pool) Submit(conn net.Conn) {
	p.connCh <- conn
}

func (p *Pool) Shutdown() {
	//Close before wait to prevent deadlock between channel closing and workers finishing tasks
	close(p.connCh)
	p.wg.Wait()
}

func (p *Pool) handleConn(conn net.Conn) {
	defer conn.Close()

	request, err := Parser(conn)
	if err != nil {
		response := "HTTP/1.1 400 Bad Request\r\n\r\n"
		conn.Write([]byte(response))
		return
	}

	handlerFunc, _ := p.router.Match(request)
	if handlerFunc == nil {
		response := "HTTP/1.1 404 Not Found\r\n\r\n"
		conn.Write([]byte(response))
		return
	}

	_ = handlerFunc(request)
	response := "HTTP/1.1 200 OK \r\n\r\n"
	conn.Write([]byte(response))
}
