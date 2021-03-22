package ssdb

import (
	"bufio"
	"io"
	"net"
	"strconv"
	"time"
)

// Conn a SSDB connection
type Conn struct {
	conn             net.Conn
	r                *bufio.Reader
	w                *bufio.Writer
	readWriteTimeout time.Duration
}

func newConn(conn net.Conn, readWriteTimeout time.Duration) *Conn {
	return &Conn{
		conn:             conn,
		r:                bufio.NewReader(conn),
		w:                bufio.NewWriter(conn),
		readWriteTimeout: readWriteTimeout,
	}
}

var zeroTime = time.Time{}

// Send send data
func (c *Conn) Send(args Values) error {
	if c.readWriteTimeout != 0 {
		if err := c.conn.SetWriteDeadline(time.Now().Add(c.readWriteTimeout)); err == nil {
			defer func() {
				_ = c.conn.SetWriteDeadline(zeroTime)
			}()
		}
	}

	for _, arg := range args {
		c.w.Write(strconv.AppendInt(nil, int64(len(arg)), 10))
		c.w.WriteByte('\n')
		c.w.Write(arg)
		c.w.WriteByte('\n')
	}
	c.w.WriteByte('\n')
	return c.w.Flush()
}

// Recv receive	data
func (c *Conn) Recv() (Values, error) {
	if c.readWriteTimeout != 0 {
		if err := c.conn.SetWriteDeadline(time.Now().Add(c.readWriteTimeout)); err == nil {
			defer func() {
				_ = c.conn.SetWriteDeadline(zeroTime)
			}()
		}
	}

	resp := Values{}
loop:
	for {
		tmp, err := c.r.ReadSlice('\n')
		if err != nil {
			return nil, err
		}

		switch len(tmp) {
		case 0:
			continue loop
		case 2:
			if tmp[0] == '\r' {
				if len(resp) == 0 {
					continue loop
				}
				return resp, nil
			}
		case 1:
			if len(resp) == 0 {
				continue loop
			}
			return resp, nil
		}

		size, err := strconv.ParseInt(string(tmp[:len(tmp)-1]), 0, 0)
		if err != nil || size < 0 {
			return nil, err
		}
		buf := make([]byte, size)
		_, err = io.ReadFull(c.r, buf)
		if err != nil {
			return nil, err
		}
		resp = append(resp, Value(buf))
		c.r.ReadByte()
	}
}

// Close Connection
func (c *Conn) Close() error {
	return c.conn.Close()
}
