package ssdb

import (
	"net"
	"net/url"
)

type Option func(c *Client)

// Url
// ssdb://127.0.0.1:8888[?Auth=password]
func Url(u string) Option {
	uu, _ := url.Parse(u)
	return func(c *Client) {
		c.addr = uu.Host
		c.auth = uu.Query().Get("Auth")
	}
}

// Addr
// 127.0.0.1:8888
func Addr(addr string) Option {
	return func(c *Client) {
		c.addr = addr
	}
}

// Auth password
func Auth(auth string) Option {
	return func(c *Client) {
		c.auth = auth
	}
}

// DialHandler proxy
func DialHandler(df func(addr string) (net.Conn, error)) Option {
	return func(c *Client) {
		c.dialHandler = df
	}
}