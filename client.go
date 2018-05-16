package ssdb

import (
	"bytes"
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
	buf  bytes.Buffer
}

// ConnectByConn Single connected client by net.Conn
func ConnectByConn(conn net.Conn) (*Client, error) {
	return &Client{
		sock: conn,
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

func (c *Client) doMapStringInt(args ...interface{}) (map[string]int64, error) {
	v, err := c.Do(args...)
	if err != nil {
		return nil, makeError(err, v, args)
	}
	if !(len(v) > 0 && v[0].Equal(ok)) {
		return nil, makeError(nil, v, args)
	}
	return v[1:].MapStringInt(), nil
}

// Send msg
func (c *Client) Send(args ...interface{}) error {
	return c.send(args)
}

func (c *Client) send(args []interface{}) error {
	var buf bytes.Buffer
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
		case []string:
			for _, s := range arg {
				buf.WriteString(fmt.Sprintf("%d", len(s)))
				buf.WriteByte('\n')
				buf.WriteString(s)
				buf.WriteByte('\n')
			}
			continue
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
		buf.WriteString(fmt.Sprintf("%d", len(s)))
		buf.WriteByte('\n')
		buf.WriteString(s)
		buf.WriteByte('\n')
	}
	buf.WriteByte('\n')
	_, err := c.sock.Write(buf.Bytes())
	return err
}

// Recv msg
func (c *Client) Recv() (Values, error) {
	var tmp [8192]byte
	for {
		resp := c.parse()
		if resp != nil {
			return resp, nil
		}
		n, err := c.sock.Read(tmp[0:])
		if err != nil {
			return nil, err
		}
		c.buf.Write(tmp[0:n])
	}
}

func (c *Client) parse() Values {
	resp := Values{}
	buf := c.buf.Bytes()
	var idx, offset int
	idx = 0
	offset = 0

	for {
		idx = bytes.IndexByte(buf[offset:], '\n')
		if idx == -1 {
			break
		}
		p := buf[offset : offset+idx]
		offset += idx + 1

		if len(p) == 0 || (len(p) == 1 && p[0] == '\r') {
			if len(resp) == 0 {
				continue
			}
			c.buf.Reset()
			c.buf.Write(buf[offset:])
			return resp
		}

		size, err := strconv.Atoi(string(p))
		if err != nil || size < 0 {
			break
		}
		if offset+size >= c.buf.Len() {
			break
		}

		v := buf[offset : offset+size]
		resp = append(resp, Value(v))
		offset += size + 1
	}

	return nil
}

// Close Connection
func (c *Client) Close() error {
	return c.sock.Close()
}
