package ssdb

import (
	"net"
	"net/url"
	"time"
)

// Option is a function that configures a Client
type Option func(c *Client)

// URL ssdb://127.0.0.1:8888[?Auth=password]
func URL(u string) Option {
	uu, _ := url.Parse(u)
	return func(c *Client) {
		c.addr = uu.Host
		c.auth = uu.Query().Get("Auth")
	}
}

// Addr 127.0.0.1:8888
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

func DialTimeoutOption(connectTimeout time.Duration) Option {
	return func(c *Client) {
		c.dialHandler = func(addr string) (conn net.Conn, e error) {
			return net.DialTimeout("tcp", addr, connectTimeout)
		}
	}
}

func ReadWriteTimeoutOption(readWriteTimeout time.Duration) Option {
	return func(c *Client) {
		c.readWriteTimeout = readWriteTimeout
	}
}
func IgnoreGetNotFoundError(ignoreGetNotFoundError bool) Option {
	return func(c *Client) {
		c.ignoreGetNotFoundError = ignoreGetNotFoundError
	}
}
