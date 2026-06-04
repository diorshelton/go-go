package server

import (
	"net"
	"sync"
	"time"
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

	for {
		conn.SetDeadline(time.Now().Add(30 * time.Second))

		request, err := Parse(conn)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				break
			}
			response := Response{StatusCode: 400, Body: nil}
			conn.Write(response.Serialize())
			break
		}

		handlerFunc, params := p.router.Match(request)
		if handlerFunc == nil {
			response := Response{StatusCode: 404, Body: nil}
			_, writeErr := conn.Write(response.Serialize())
			if writeErr != nil {
				break
			}
			continue
		}

		request.Params = params

		response := handlerFunc(request)

		_, writeErr := conn.Write(response.Serialize())
		if writeErr != nil {
			break
		}

		if request.Headers["connection"] == "close" {
			break
		}

	}
}
