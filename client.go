package ssdb

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

var ok = Value("ok")
var notFound = Value("not_found")
var one = Value("1")

// Client Single connected client
type Client struct {
	sock net.Conn
	rw   *bufio.ReadWriter
}

// ConnectByConn Single connected client by net.Conn
func ConnectByConn(conn net.Conn) (*Client, error) {
	return &Client{
		rw: bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)),
	}, nil
}

// ConnectByAddr Single connected client by addr
func ConnectByAddr(addr string) (*Client, error) {
	sock, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return ConnectByConn(sock)
}

// Send msg
func (c *Client) Send(args ...interface{}) error {
	return c.send(args)
}

func (c *Client) send(args []interface{}) error {
	for _, arg := range args {
		var s string
		switch arg := arg.(type) {
		case time.Duration:
			s = strconv.FormatUint(uint64(arg/time.Second), 10)
		case fmt.Stringer:
			s = arg.String()
		case string:
			s = arg
		case []byte:
			s = string(arg)
		case int:
			s = strconv.FormatInt(int64(arg), 10)
		case int8:
			s = strconv.FormatInt(int64(arg), 10)
		case int16:
			s = strconv.FormatInt(int64(arg), 10)
		case int32:
			s = strconv.FormatInt(int64(arg), 10)
		case int64:
			s = strconv.FormatInt(int64(arg), 10)
		case uint:
			s = strconv.FormatUint(uint64(arg), 10)
		case uint8:
			s = strconv.FormatUint(uint64(arg), 10)
		case uint16:
			s = strconv.FormatUint(uint64(arg), 10)
		case uint32:
			s = strconv.FormatUint(uint64(arg), 10)
		case uint64:
			s = strconv.FormatUint(uint64(arg), 10)
		case float32:
			s = strconv.FormatFloat(float64(arg), 'f', -1, 64)
		case float64:
			s = strconv.FormatFloat(float64(arg), 'f', -1, 64)
		case bool:
			if arg {
				s = "1"
			} else {
				s = "0"
			}
		case nil:
			s = ""
		default:
			return fmt.Errorf("bad arguments")
		}

		c.rw.WriteString(fmt.Sprintf("%d\n", len(s)))
		c.rw.WriteString(s)
		c.rw.WriteByte('\n')
	}
	c.rw.WriteByte('\n')
	return c.rw.Flush()
}

// Recv msg
func (c *Client) Recv() (Values, error) {
	return c.recv()
}

func (c *Client) recv() (Values, error) {
	resp := Values{}
	for {
		tmp, err := c.rw.ReadSlice('\n')
		if err != nil {
			return nil, err
		}

		if len(tmp) == 0 {
			continue
		}

		if tmp[0] == '\n' || tmp[0] == '\r' {
			if len(resp) == 0 {
				continue
			}
			return resp, nil
		}
		size, err := strconv.Atoi(string(tmp[:len(tmp)-1]))
		if err != nil || size < 0 {
			return nil, err
		}
		buf := make([]byte, size)
		_, err = c.rw.Read(buf)
		if err != nil {
			return nil, err
		}
		resp = append(resp, Value(buf))
		c.rw.ReadByte()
	}
}

// Close Connection
func (c *Client) Close() error {
	return c.sock.Close()
}
