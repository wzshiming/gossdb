package ssdb

import (
	"net"
	"sync"
)

var (
	ok          = Value("ok")
	notFound    = Value("not_found")
	clientError = Value("client_error")
	one         = Value("1")
	zero        = Value("0")
)

// Client connected client
type Client struct {
	pool        sync.Pool
	dialHandler func(addr string) (net.Conn, error)
	auth        string
	addr        string
}

// Connect connected client
func Connect(opts ...Option) (*Client, error) {
	c := &Client{
		dialHandler: func(addr string) (net.Conn, error) {
			return net.Dial("tcp", addr)
		},
		addr: "127.0.0.1:8888",
	}
	c.pool = sync.Pool{
		New: func() interface{} {
			netConn, err := c.dialHandler(c.addr)
			if err != nil {
				return err
			}
			conn := newConn(netConn)

			err = conn.Send(Values{Value("auth"), Value(c.auth)})
			if err != nil {
				return err
			}
			_, err = conn.Recv()
			if err != nil {
				return err
			}
			return conn
		},
	}
	for _, v := range opts {
		v(c)
	}

	return c, nil
}
