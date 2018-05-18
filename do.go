package ssdb

import (
	"fmt"
	"time"
)

// GetConn get connection
func (c *Client) GetConn() (*Conn, error) {
	conni := c.pool.Get()
	switch t := conni.(type) {
	case *Conn:
		return t, nil
	case error:
		return nil, t
	default:
		return nil, fmt.Errorf("Error version")
	}
}

// PutConn Put back the connection
func (c *Client) PutConn(conn *Conn) {
	if conn == nil {
		return
	}
	c.pool.Put(conn)
}

// Do send and recv
func (c *Client) Do(args ...interface{}) (v Values, err0 error) {
	conn, err := c.GetConn()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err0 != nil {
			c.pool.Put(conn)
		} else {
			conn.Close()
		}
	}()
	v, err = NewValues(args)
	if err != nil {
		return nil, err
	}
	err = conn.Send(v)
	if err != nil {
		return nil, err
	}
	return conn.Recv()
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

func (c *Client) doMapStringValue(args ...interface{}) (map[string]Value, error) {
	v, err := c.Do(args...)
	if err != nil {
		return nil, makeError(err, v, args)
	}
	if !(len(v) > 0 && v[0].Equal(ok)) {
		return nil, makeError(nil, v, args)
	}
	return v[1:].MapStringValue(), nil
}

func (c *Client) doDuration(args ...interface{}) (time.Duration, error) {
	v, err := c.Do(args...)
	if err != nil {
		return -1, makeError(err, v, args)
	}
	if !(len(v) > 1 && v[0].Equal(ok)) {
		return -1, makeError(nil, v, args)
	}
	return v[1].Duration(), nil
}

func (c *Client) doInt(args ...interface{}) (int64, error) {
	v, err := c.Do(args...)
	if err != nil {
		return -1, makeError(err, v, args)
	}
	if !(len(v) > 1 && v[0].Equal(ok)) {
		return -1, makeError(nil, v, args)
	}
	return v[1].Int(), nil
}

func (c *Client) doBool(args ...interface{}) (bool, error) {
	v, err := c.Do(args...)
	if err != nil {
		return false, makeError(err, v, args)
	}
	if !(len(v) > 1 && v[0].Equal(ok)) {
		return false, makeError(nil, v, args)
	}
	return v[1].Equal(one), nil
}

func (c *Client) doString(args ...interface{}) (string, error) {
	v, err := c.Do(args...)
	if err != nil {
		return "", makeError(err, v, args)
	}
	if !(len(v) > 1 && v[0].Equal(ok)) {
		return "", makeError(nil, v, args)
	}
	return v[1].String(), nil
}

func (c *Client) doStrings(args ...interface{}) ([]string, error) {
	v, err := c.Do(args...)
	if err != nil {
		return nil, makeError(err, v, args)
	}
	if !(len(v) > 0 && v[0].Equal(ok)) {
		return nil, makeError(nil, v, args)
	}
	return v[1:].Strings(), nil
}

func (c *Client) doValue(args ...interface{}) (Value, error) {
	v, err := c.Do(args...)
	if err != nil {
		return nil, makeError(err, v, args)
	}
	if !(len(v) > 1 && v[0].Equal(ok)) {
		return nil, makeError(nil, v, args)
	}
	return v[1], nil
}

func (c *Client) doValues(args ...interface{}) (Values, error) {
	v, err := c.Do(args...)
	if err != nil {
		return nil, makeError(err, v, args)
	}
	if !(len(v) > 0 && v[0].Equal(ok)) {
		return nil, makeError(nil, v, args)
	}
	return v[1:], nil
}

func (c *Client) doNil(args ...interface{}) error {
	v, err := c.Do(args...)
	if err != nil {
		return makeError(err, v, args)
	}
	if !(len(v) > 0 && v[0].Equal(ok)) {
		return makeError(nil, v, args)
	}
	return nil
}

func (c *Client) doCDStringValue(cb func(string, Value) error, startPos int, limit int64, args ...interface{}) error {
	vs, err := c.doValues(args...)
	if err != nil {
		return err
	}

	var k string
	for i, end := 0, len(vs)-1; i < end; i += 2 {
		k = vs[i].String()
		v := vs[i+1]
		err = cb(k, v)
		if err != nil {
			return err
		}
	}
	if int64(len(vs)/2) == limit {
		args[startPos] = k
		return c.doCDStringValue(cb, startPos, limit, args...)
	}
	return nil
}

func (c *Client) doCDString(cb func(string) error, startPos int, limit int64, args ...interface{}) error {
	vs, err := c.doStrings(args...)
	if err != nil {
		return err
	}

	for _, v := range vs {
		err = cb(v)
		if err != nil {
			return err
		}
	}

	if int64(len(vs)) == limit {
		args[startPos] = vs[len(vs)-1]
		return c.doCDString(cb, startPos, limit, args...)
	}
	return nil
}
