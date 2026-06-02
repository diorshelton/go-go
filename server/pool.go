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

	request, err := Parse(conn)
	if err != nil {
		response := Response{StatusCode: 400, Body: nil}
		conn.Write(response.Serialize())
		return
	}

	handlerFunc, params := p.router.Match(request)
	if handlerFunc == nil {
		response := Response{StatusCode: 404, Body: nil}
		conn.Write(response.Serialize())
		return
	}

	request.Params = params

	responseReturn := handlerFunc(request)

	response := responseReturn.Serialize()

	conn.Write(response)
}
