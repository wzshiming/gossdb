package gossdb

// HSet name key value
// Set the string value in argument as value of the key of a hashmap.
func (c *Client) HSet(name, key string, value interface{}) (bool, error) {
	return c.doBool("hset", name, key, value)
}

// HGet name key
// Get the value related to the specified key of a hashmap
func (c *Client) HGet(name, key string) (Value, error) {
	return c.doValue("hget", name, key)
}

// HDel name key
// Delete specified key of a hashmap. To delete the whole hashmap, use hclear.
func (c *Client) HDel(name, key string) (bool, error) {
	return c.doBool("hdel", name, key)
}

// HIncr name key [num]
// Since 1.7.0.1, *incr methods return error if value cannot be converted to integer.
// Increment the number stored at key in a hashmap by num. The num argument could be a negative integer. The old number is first converted to an integer before increment, assuming it was stored as literal integer.
func (c *Client) HIncr(name, key string, num int64) (int64, error) {
	return c.doInt("hincr", name, key, num)
}

// HExists name key
// Verify if the specified key exists in a hashmap.
func (c *Client) HExists(name, key string) (bool, error) {
	return c.doBool("hexists", name, key)
}

// HSize name
// Return the number of key-value pairs in the hashmap.
func (c *Client) HSize(name string) (int64, error) {
	return c.doInt("hsize", name)
}

// HList nameStart nameEnd limit
// List hashmap names in range (nameStart, nameEnd].
// ("", ""] means no range limit.
// Refer to scan command for more information about how it work.
func (c *Client) HList(nameStart, nameEnd string, limit int64) (Values, error) {
	return c.doValues("hlist", nameStart, nameEnd, limit)
}

// HRList nameStart nameEnd limit
// Like hlist, but in reverse order.
func (c *Client) HRList(nameStart, nameEnd string, limit int64) (Values, error) {
	return c.doValues("hrlist", nameStart, nameEnd, limit)
}

// HKeys name keyStart keyEnd limit
// List keys of a hashmap in range (keyStart, keyEnd].
// ("", ""] means no range limit.
func (c *Client) HKeys(name, keyStart, keyEnd string, limit int64) (Values, error) {
	return c.doValues("hkeys", name, keyStart, keyEnd, limit)
}

// HGetAll name
// Returns the whole hash, as an array of strings indexed by strings.
func (c *Client) HGetAll(name string) (map[string]Value, error) {
	return c.doMapStringValue("hgetall", name)
}

// HScan name keyStart keyEnd limit
// List key-value pairs of a hashmap with keys in range (keyStart, keyEnd].
// ("", ""] means no range limit.
// Refer to scan command for more information about how it work.
func (c *Client) HScan(name string, keyStart, keyEnd string, limit int64) (map[string]Value, error) {
	return c.doMapStringValue("hscan", name, keyStart, keyEnd, limit)
}

// HRScan name keyStart keyEnd limit
// Like hscan, but in reverse order.
func (c *Client) HRScan(name string, keyStart, keyEnd string, limit int64) (map[string]Value, error) {
	return c.doMapStringValue("hrscan", name, keyStart, keyEnd, limit)
}

// HClear name
// Delete all keys in a hashmap.
func (c *Client) HClear(name string) error {
	return c.doNil("hclear", name)
}

// MultiHSet name key1 value1 key2 value2 ...
// Set multiple key-value pairs(kvs) of a hashmap in one method call.
func (c *Client) MultiHSet(name string, kvs map[string]interface{}) error {

	args := []interface{}{"multi_hset", name}
	for k, v := range kvs {
		args = append(args, k)
		args = append(args, v)
	}
	return c.doNil(args...)
}

// MultiHGet name key1 key2 ...
// Get the values related to the specified multiple keys of a hashmap.
func (c *Client) MultiHGet(name string, key ...string) (map[string]Value, error) {
	if len(key) == 0 {
		return make(map[string]Value), nil
	}

	args := []interface{}{"multi_hget", name}

	for _, v := range key {
		args = append(args, v)
	}

	return c.doMapStringValue(args...)
}

// MultiHDel name key1 key2 ...
// Delete specified multiple keys in a hashmap.
func (c *Client) MultiHDel(name string, key ...string) error {
	if len(key) == 0 {
		return nil
	}
	args := []interface{}{"multi_hdel", name}
	for _, v := range key {
		args = append(args, v)
	}
	return c.doNil(args...)
}
