package ssdb

import "time"

// Do send and recv
func (c *Client) Do(args ...interface{}) (Values, error) {
	err := c.send(args)
	if err != nil {
		return nil, err
	}
	resp, err := c.Recv()
	return resp, err
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
	if !(len(v) > 1 && v[0].Equal(ok)) {
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
	if !(len(v) > 1 && v[0].Equal(ok)) {
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
