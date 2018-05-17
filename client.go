package ssdb

import (
	"net"
	"sync"
)

var ok = Value("ok")
var notFound = Value("not_found")
var one = Value("1")

// Client Single connected client
type Client struct {
	pool        sync.Pool
	dialHandler func(addr string) (net.Conn, error)
	auth        string
	addr        string
}

// Connect Single connected client by net.Conn
func Connect(opts ...Option) (*Client, error) {
	c := &Client{
		dialHandler: func(addr string) (net.Conn, error) {
			return net.Dial("tcp", addr)
		},
		addr: "127.0.0.1:8888",
	}
	c.pool = sync.Pool{
		New: func() interface{} {
			conn, err := c.dialHandler(c.addr)
			if err != nil {
				return err
			}
			return newConn(conn)
		},
	}
	for _, v := range opts {
		v(c)
	}

	return c, nil
}
