package ssdb

import "net"

type Option func(c *Client)

func Addr(addr string) Option {
	return func(c *Client) {
		c.addr = addr
	}
}

func Auth(auth string) Option {
	return func(c *Client) {
		c.auth = auth
	}
}

func DialHandler(df func(addr string) (net.Conn, error)) Option {
	return func(c *Client) {
		c.dialHandler = df
	}
}
