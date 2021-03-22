package ssdb

// ZSet Set the score of the key of a zset.
func (c *Client) ZSet(name, key string, score int64) error {
	return c.doNil("zset", name, key, score)
}

// ZGet Get the score related to the specified key of a zset
func (c *Client) ZGet(name, key string) (int64, error) {
	return c.doInt("zget", name, key)
}

// ZDel Delete specified key of a zset.
func (c *Client) ZDel(name, key string) error {
	return c.doNil("zdel", name, key)
}

// ZIncr Increment the number stored at key in a zset by num.
func (c *Client) ZIncr(name string, key string, num int64) (int64, error) {
	return c.doInt("zincr", name, key, num)
}

// ZExists Verify if the specified key exists in a zset.
func (c *Client) ZExists(name, key string) (bool, error) {
	return c.doBool("zexists", name, key)
}

// ZSize Return the number of pairs of a zset.
func (c *Client) ZSize(name string) (int64, error) {
	return c.doInt("zsize", name)
}

// ZList List zset names in range (nameStart, nameEnd].
// Refer to scan command for more information about how it work.
func (c *Client) ZList(nameStart, nameEnd string, limit int64) ([]string, error) {
	return c.doStrings("zlist", nameStart, nameEnd, limit)
}

// ZListRangeAll Like zlist, The whole range
func (c *Client) ZListRangeAll(nameStart, nameEnd string, limit int64, cb func(string) error) error {
	return c.doCDString(cb, 1, limit, "zlist", nameStart, nameEnd, limit)
}

// ZRList List zset names in range (nameStart, nameEnd], in reverse order.
func (c *Client) ZRList(nameStart, nameEnd string, limit int64) ([]string, error) {
	return c.doStrings("zrlist", nameStart, nameEnd, limit)
}

// ZRListRangeAll Like zrlist, The whole range
func (c *Client) ZRListRangeAll(nameStart, nameEnd string, limit int64, cb func(string) error) error {
	return c.doCDString(cb, 1, limit, "zrlist", nameStart, nameEnd, limit)
}

// ZKeys List keys in a zset.
func (c *Client) ZKeys(name string, keyStart string, scoreStart, scoreEnd int64, limit int64) ([]string, error) {
	return c.doStrings("zkeys", name, keyStart, scoreStart, scoreEnd, limit)
}

// ZKeysRangeAll Like zkeys, The whole range
func (c *Client) ZKeysRangeAll(name string, keyStart string, scoreStart, scoreEnd int64, limit int64, cb func(string) error) error {
	return c.doCDString(cb, 2, limit, "zkeys", name, keyStart, scoreStart, scoreEnd, limit)
}

// ZScan List key-score pairs where key-score in range (keyStart+scoreStart, scoreEnd].
// Refer to scan command for more information about how it work.
// map: ZScan -> ZScanMap
func (c *Client) ZScan(name string, keyStart string, scoreStart, scoreEnd int64, limit int64) (map[string]int64, error) {
	return c.ZScanMap(name, keyStart, scoreStart, scoreEnd, limit)
}

// ZScanMap (ZScan) List key-score pairs where key-score in range (keyStart+scoreStart, scoreEnd].
// Refer to scan command for more information about how it work.
func (c *Client) ZScanMap(name string, keyStart string, scoreStart, scoreEnd int64, limit int64) (map[string]int64, error) {
	return c.doMapStringInt("zscan", name, keyStart, scoreStart, scoreEnd, limit)
}

// ZScanSlice (ZScan) List key-score pairs where key-score in range (keyStart+scoreStart, scoreEnd].
// Refer to scan command for more information about how it work.
// This command was keep elements sort to ZScan
func (c *Client) ZScanSlice(name string, keyStart string, scoreStart, scoreEnd int64, limit int64) (ValueScoreSlice, error) {
	return c.doValueScoreSlice("zscan", name, keyStart, scoreStart, scoreEnd, limit)
}

// ZRScan List key-score pairs of a zset, in reverse order.
// Deprecated: ZRScan -> ZRScanMap
func (c *Client) ZRScan(name string, keyStart string, scoreStart, scoreEnd int64, limit int64) (map[string]int64, error) {
	return c.ZRScanMap(name, keyStart, scoreStart, scoreEnd, limit)
}

// ZRScanMap (ZRScan) List key-score pairs of a zset, in reverse order.
func (c *Client) ZRScanMap(name string, keyStart string, scoreStart, scoreEnd int64, limit int64) (map[string]int64, error) {
	return c.doMapStringInt("zrscan", name, keyStart, scoreStart, scoreEnd, limit)
}

// ZRScanSlice (ZRScan) List key-score pairs of a zset, in reverse order.
// This command was keep elements sort to ZRScan
func (c *Client) ZRScanSlice(name string, keyStart string, scoreStart, scoreEnd int64, limit int64) (ValueScoreSlice, error) {
	return c.doValueScoreSlice("zrscan", name, keyStart, scoreStart, scoreEnd, limit)
}

// ZRank Returns the rank(index) of a given key in the specified sorted set.
func (c *Client) ZRank(name, key string) (int64, error) {
	return c.doInt("zrank", name, key)
}

// ZRRank Returns the rank(index) of a given key in the specified sorted set, in reverse order.
func (c *Client) ZRRank(name, key string) (int64, error) {
	return c.doInt("zrrank", name, key)
}

// ZRange Returns a range of key-score pairs by index range [offset, offset + limit).
// Deprecated: ZRange -> ZRangeMap
func (c *Client) ZRange(name string, offset, limit int64) (map[string]int64, error) {
	return c.ZRangeMap(name, offset, limit)
}

// ZRangeMap (ZRange) Returns a range of key-score pairs by index range [offset, offset + limit).
func (c *Client) ZRangeMap(name string, offset, limit int64) (map[string]int64, error) {
	return c.doMapStringInt("zrange", name, offset, limit)
}

// ZRangeSlice (ZRange) Returns a range of key-score pairs by index range [offset, offset + limit).
// This command was keep elements sort to ZRange
func (c *Client) ZRangeSlice(name string, offset, limit int64) (ValueScoreSlice, error) {
	return c.doValueScoreSlice("zrange", name, offset, limit)
}

// ZRRange Returns a range of key-score pairs by index range [offset, offset + limit), in reverse order.
// Deprecated: ZRRange => ZRRangeMap
func (c *Client) ZRRange(name string, offset, limit int64) (map[string]int64, error) {
	return c.ZRRangeMap(name, offset, limit)
}

// ZRRangeMap (ZRRange) Returns a range of key-score pairs by index range [offset, offset + limit), in reverse order.
func (c *Client) ZRRangeMap(name string, offset, limit int64) (map[string]int64, error) {
	return c.doMapStringInt("zrrange", name, offset, limit)
}

// ZRRangeSlice (ZRRange) Returns a range of key-score pairs by index range [offset, offset + limit), in reverse order.
// This command was keep elements sort to ZRRange
func (c *Client) ZRRangeSlice(name string, offset, limit int64) (ValueScoreSlice, error) {
	return c.doValueScoreSlice("zrrange", name, offset, limit)
}

// ZClear Delete all keys in a zset.
func (c *Client) ZClear(name string) error {
	return c.doNil("zclear", name)
}

// ZCount Returns the number of elements of the sorted set stored at the specified key which have scores in the range [start,end].
func (c *Client) ZCount(name string, start, end string) (int64, error) {
	return c.doInt("zcount", name, start, end)
}

// ZSum Returns the sum of elements of the sorted set stored at the specified key which have scores in the range [start,end].
func (c *Client) ZSum(name string, scoreStart, scoreEnd int64) (int64, error) {
	return c.doInt("zsum", name, scoreStart, scoreEnd)
}

// ZAvg Returns the average of elements of the sorted set stored at the specified key which have scores in the range [start,end].
func (c *Client) ZAvg(name string, scoreStart, scoreEnd int64) (int64, error) {
	return c.doInt("zavg", name, scoreStart, scoreEnd)
}

// ZRemRangeByRank Delete the elements of the zset which have rank in the range [start,end].
func (c *Client) ZRemRangeByRank(name string, start, end int64) error {
	return c.doNil("zremrangebyrank", name, start, end)
}

// ZRemRangeByScore Delete the elements of the zset which have score in the range [start,end].
func (c *Client) ZRemRangeByScore(name string, start, end int64) error {
	return c.doNil("zremrangebyscore", name, start, end)
}

// ZPopFront Since 1.9.0, Delete and return limit element(s) from front of the zset.
// Deprecated: ZPopFront -> ZPopFrontMap
func (c *Client) ZPopFront(name string, limit int64) (map[string]int64, error) {
	return c.ZPopFrontMap(name, limit)
}

// ZPopFrontMap (ZPopFront) Since 1.9.0, Delete and return limit element(s) from front of the zset.
func (c *Client) ZPopFrontMap(name string, limit int64) (map[string]int64, error) {
	return c.doMapStringInt("zpop_front", name, limit)
}

// ZPopFrontSlice (ZPopFront) Since 1.9.0, Delete and return limit element(s) from front of the zset.
// This command was keep elements sort to ZPopFront
func (c *Client) ZPopFrontSlice(name string, limit int64) (ValueScoreSlice, error) {
	return c.doValueScoreSlice("zpop_front", name, limit)
}

// ZPopBack Since 1.9.0, Delete and return limit element(s) from back of the zset.
// Deprecated: ZPopBack -> ZPopBackMap
func (c *Client) ZPopBack(name string, limit int64) (map[string]int64, error) {
	return c.ZPopBackMap(name, limit)
}

// ZPopBackMap (ZPopBack) Since 1.9.0, Delete and return limit element(s) from back of the zset.
func (c *Client) ZPopBackMap(name string, limit int64) (map[string]int64, error) {
	return c.doMapStringInt("zpop_back", name, limit)
}

// ZPopBackSlice (ZPopBack) Since 1.9.0, Delete and return limit element(s) from back of the zset.
// This command was keep elements sort to ZPopBack
func (c *Client) ZPopBackSlice(name string, limit int64) (ValueScoreSlice, error) {
	return c.doValueScoreSlice("zpop_back", name, limit)
}

// MultiZSet Set multiple key-score pairs(kvs) of a zset in one method call.
func (c *Client) MultiZSet(name string, kvs map[string]int64) error {
	args := []interface{}{"multi_zset", name}
	for k, v := range kvs {
		args = append(args, k)
		args = append(args, v)
	}
	return c.doNil(args...)
}

// MultiZGet Get the values related to the specified multiple keys of a zset.
func (c *Client) MultiZGet(name string, key ...string) (map[string]int64, error) {
	if len(key) == 0 {
		return map[string]int64{}, nil
	}
	args := []interface{}{"multi_zget", name}

	for _, v := range key {
		args = append(args, v)
	}

	return c.doMapStringInt(args...)
}

// MultiZDel Delete specified multiple keys of a zset.
func (c *Client) MultiZDel(name string, key ...string) error {
	if len(key) == 0 {
		return nil
	}
	args := []interface{}{"multi_zdel", name}
	for _, v := range key {
		args = append(args, v)
	}
	return c.doNil(args...)
}
