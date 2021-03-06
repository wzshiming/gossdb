package ssdb

// QPushFront Add one or more than one element to the head of the queue.
func (c *Client) QPushFront(name string, item ...Value) (int64, error) {
	args := []interface{}{"qpush_front", name}
	for _, v := range item {
		args = append(args, v)
	}
	return c.doInt(args...)
}

// QPushBack Add an or more than one element to the end of the queue.
func (c *Client) QPushBack(name string, item ...Value) (int64, error) {
	args := []interface{}{"qpush_back", name}
	for _, v := range item {
		args = append(args, v)
	}
	return c.doInt(args...)
}

// QPopFront Pop out one or more elements from the head of a queue.
func (c *Client) QPopFront(name string, size int64) (Values, error) {
	return c.doValues("qpop_front", name, size)
}

// QPopBack Pop out one or more elements from the tail of a queue.
func (c *Client) QPopBack(name string, size int64) (Values, error) {
	return c.doValues("qpop_back", name, size)
}

// QFront Returns the first element of a queue.
func (c *Client) QFront(key string) (Value, error) {
	return c.doValue("qfront", key)
}

// QBack Returns the last element of a queue.
func (c *Client) QBack(key string) (Value, error) {
	return c.doValue("qback", key)
}

// QSize Returns the number of items in the queue.
func (c *Client) QSize(name string) (int64, error) {
	return c.doInt("qsize", name)
}

// QClear Clear the queue.
func (c *Client) QClear(name string) error {
	return c.doNil("qclear", name)
}

// QGet Returns the element a the specified index(position).
// 0 the first element, 1 the second ... -1 the last element.
func (c *Client) QGet(key string, index int64) (Value, error) {
	return c.doValue("qget", key, index)
}

// QSet Sets the list element at index to value.
// An error is returned for out of range indexes.
func (c *Client) QSet(key string, index int64, value Value) error {
	return c.doNil("qset", key, index, value)
}

// QRange Returns a portion of elements from the queue at the specified range [offset, offset + limit].
func (c *Client) QRange(name string, offset, limit int) (Values, error) {
	return c.doValues("qrange", name, offset, limit)
}

// QSlice Returns a portion of elements from the queue at the specified range [begin, end].
// begin and end could be negative.
func (c *Client) QSlice(name string, begin, end int) (Values, error) {
	return c.doValues("qslice", name, begin, end)
}

// QTrimFront Remove multi elements from the head of a queue.
func (c *Client) QTrimFront(name string, size int) (int64, error) {
	return c.doInt("qtrim_front", name, size)
}

// QTrimBack Remove multi elements from the tail of a queue.
func (c *Client) QTrimBack(name string, size int) (int64, error) {
	return c.doInt("qtrim_back", name, size)
}

// QList List list/queue names in range (nameStart, nameEnd].
// ("", ""] means no range limit.
// Refer to scan command for more information about how it work.
func (c *Client) QList(nameStart, nameEnd string, limit int64) ([]string, error) {
	return c.doStrings("qlist", nameStart, nameEnd, limit)
}

// QListRangeAll Like qlist, The whole range
func (c *Client) QListRangeAll(nameStart, nameEnd string, limit int64, cb func(string) error) error {
	return c.doCDString(cb, 1, limit, "qlist", nameStart, nameEnd, limit)
}

// QRList Like qlist, but in reverse order.
func (c *Client) QRList(nameStart, nameEnd string, limit int64) ([]string, error) {
	return c.doStrings("qrlist", nameStart, nameEnd, limit)
}

// QRListRangeAll Like qrlist, The whole range
func (c *Client) QRListRangeAll(nameStart, nameEnd string, limit int64, cb func(string) error) error {
	return c.doCDString(cb, 1, limit, "qrlist", nameStart, nameEnd, limit)
}
