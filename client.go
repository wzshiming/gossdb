package ssdb

import (
	"net"
)

var ok = Value("ok")
var notFound = Value("not_found")
var one = Value("1")

// Client Single connected client
type Client struct {
	conn *Conn
}

// Connect Single connected client by net.Conn
func Connect(f func() (net.Conn, error)) (*Client, error) {
	conn, err := f()
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: newConn(conn),
	}, nil
}

// ConnectByAddr Single connected client by addr
func ConnectByAddr(addr string) (*Client, error) {
	return Connect(func() (net.Conn, error) {
		return net.Dial("tcp", addr)
	})
}

// Send msg
func (c *Client) send(args []interface{}) error {
	v, err := NewValues(args)
	if err != nil {
		return err
	}
	return c.conn.Send(v)
}

// Recv msg
func (c *Client) recv() (Values, error) {
	return c.conn.Recv()
}

// Close Connection
func (c *Client) Close() error {
	return c.conn.Close()
}
