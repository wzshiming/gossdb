package ssdb

import (
	"time"
)

// Set key value
// Set Set the value of the key.
func (c *Client) Set(key string, value interface{}) error {
	return c.doNil("set", key, value)
}

// SetX key value ttl
// Set the value of the key, with a time to live.
// Unlike Redis, the ttl will not be remove when later set the same key!
func (c *Client) SetX(key string, value interface{}, ttl time.Duration) error {
	return c.doNil("setx", key, value, ttl)
}

// SetNX key value
// Set the string value in argument as value of the key if and only if the key doesn't exist.
func (c *Client) SetNX(key string, value interface{}) (bool, error) {
	return c.doBool("setnx", key, value)
}

// Expire key ttl
// Set the time left to live in seconds, only for keys of KV type.
func (c *Client) Expire(key string, ttl time.Duration) (bool, error) {
	return c.doBool("expire", key, ttl)
}

// Ttl key
// Returns the time left to live in seconds, only for keys of KV type.
func (c *Client) Ttl(key string) (time.Duration, error) {
	return c.doDuration("ttl", key)
}

// Get key
// Get the value related to the specified key.
func (c *Client) Get(key string) (Value, error) {
	return c.doValue("get", key)
}

// GetSet key value
// Sets a value and returns the previous entry at that key.
func (c *Client) GetSet(key string, value interface{}) (Value, error) {
	return c.doValue("getset", key, value)
}

// Del key
// Delete specified key.
func (c *Client) Del(key string) error {
	return c.doNil("del", key)
}

// Incr key [num]
// Since 1.7.0.1, *incr methods return error if value cannot be converted to integer.
// Increment the number stored at key by num. The num argument could be a negative integer. The old number is first converted to an integer before increment, assuming it was stored as literal integer.
func (c *Client) Incr(key string, num int64) (int64, error) {
	return c.doInt("incr", key, num)
}

// Exists key
// Verify if the specified key exists.
func (c *Client) Exists(key string) (bool, error) {
	return c.doBool("exists", key)
}

// GetBit key offset
// Return a single bit out of a string.
func (c *Client) GetBit(key string, offset int64) (bool, error) {
	return c.doBool("getbit", key, offset)
}

// SetBit key offset value
// Changes a single bit of a string. The string is auto expanded.
func (c *Client) SetBit(key string, offset int64, value bool) (bool, error) {
	return c.doBool("setbit", key, offset, value)
}

// BitCount key [start] [end]
// Count the number of set bits (population counting) in a string. Like Redis's bitcount.
func (c *Client) BitCount(key string, start, end int64) (int64, error) {
	return c.doInt("bitcount", key, start, end)
}

// CountBit key [start] [end]
// Count the number of set bits (population counting) in a string. Like Redis's bitcount.
func (c *Client) CountBit(key string, start, size int64) (int64, error) {
	return c.doInt("countbit", key, start, size)
}

// SubStr key start size
// Return part of a string,
func (c *Client) SubStr(key string, start int64, size int64) (string, error) {
	return c.doString("substr", key, start, size)
}

// StrLen key
// Return the number of bytes of a string.
func (c *Client) StrLen(key string) (int64, error) {
	return c.doInt("strlen", key)
}

// Keys keyStart keyEnd limit
// Refer to scan command for more information about how it work.
func (c *Client) Keys(keyStart, keyEnd string, limit int64) ([]string, error) {
	return c.doStrings("keys", keyStart, keyEnd, limit)
}

func (c *Client) KeysRangeAll(keyStart, keyEnd string, limit int64, cb func(string) error) error {
	return c.doCDString(cb, 1, limit, "keys", keyStart, keyEnd, limit)
}

// RKeys keyStart keyEnd limit
// Since 1.9.0
// Like keys, but in reverse order.
func (c *Client) RKeys(keyStart, keyEnd string, limit int64) ([]string, error) {
	return c.doStrings("rkeys", keyStart, keyEnd, limit)
}

func (c *Client) RKeysRangeAll(keyStart, keyEnd string, limit int64, cb func(string) error) error {
	return c.doCDString(cb, 1, limit, "rkeys", keyStart, keyEnd, limit)
}

// Scan keyStart keyEnd limit
// List key-value pairs with keys in range (keyStart, keyEnd].
// ("", ""] means no range limit.
// This command can do wildchar * like search, but only prefix search, and the * char must never occur in keyStart and keyEnd!
func (c *Client) Scan(keyStart, keyEnd string, limit int64) (map[string]Value, error) {
	return c.doMapStringValue("scan", keyStart, keyEnd, limit)
}

func (c *Client) ScanRangeAll(keyStart, keyEnd string, limit int64, cb func(string, Value) error) error {
	return c.doCDStringValue(cb, 1, limit, "scan", keyStart, keyEnd, limit)
}

// RScan
// Like scan, but in reverse order.
func (c *Client) RScan(keyStart, keyEnd string, limit int64) (map[string]Value, error) {
	return c.doMapStringValue("rscan", keyStart, keyEnd, limit)
}

func (c *Client) RScanRangeAll(keyStart, keyEnd string, limit int64, cb func(string, Value) error) error {
	return c.doCDStringValue(cb, 1, limit, "rscan", keyStart, keyEnd, limit)
}

// MultiSet key1 value1 key2 value2 ...
// Set multiple key-value pairs(kvs) in one method call.
func (c *Client) MultiSet(kvs map[string]interface{}) (err error) {
	args := []interface{}{"multi_set"}
	for k, v := range kvs {
		args = append(args, k)
		args = append(args, v)
	}
	return c.doNil(args...)
}

// MultiGet key1 key2 ...
// Get the values related to the specified multiple keys
func (c *Client) MultiGet(key ...string) (map[string]Value, error) {
	if len(key) == 0 {
		return map[string]Value{}, nil
	}
	data := []interface{}{"multi_get"}
	for _, k := range key {
		data = append(data, k)
	}
	return c.doMapStringValue(data...)
}

// MultiDel key1 key2 ...
// Delete specified multiple keys.
func (c *Client) MultiDel(key ...string) error {
	if len(key) == 0 {
		return nil
	}
	args := []interface{}{"multi_del"}
	for _, v := range key {
		args = append(args, v)
	}
	return c.doNil(args...)
}
