package gossdb

import (
	"bytes"
	"strconv"
	"time"
)

type Values []Value

func (v Values) Strings() []string {
	s := make([]string, 0, len(v))
	for _, v := range v {
		s = append(s, v.String())
	}
	return s
}

func (v Values) MapStringInt() map[string]int64 {
	val := map[string]int64{}
	size := len(v)
	for i := 0; i+1 < size; i += 2 {
		val[v[i].String()] = v[i+1].Int()
	}
	return val
}

func (v Values) MapStringValue() map[string]Value {
	val := map[string]Value{}
	size := len(v)
	for i := 0; i+1 < size; i += 2 {
		val[v[i].String()] = v[i+1]
	}
	return val
}

// Value
type Value []byte

// String
func (v Value) String() string {
	return string(v)
}

// Duration
func (v Value) Duration() time.Duration {
	return time.Duration(v.Int()) * time.Second
}

// Int
func (v Value) Int() int64 {
	i, _ := strconv.ParseInt(string(v), 0, 0)
	return i
}

// Uint
func (v Value) Uint() uint64 {
	i, _ := strconv.ParseUint(string(v), 0, 0)
	return i
}

// Float
func (v Value) Float() float64 {
	i, _ := strconv.ParseFloat(string(v), 0)
	return i
}

// Bool
func (v Value) Bool() bool {
	i, _ := strconv.ParseBool(string(v))
	return i
}

// Bytes
func (v Value) Bytes() []byte {
	return []byte(v)
}

// IsEmpty
func (v Value) IsEmpty() bool {
	return len(v) == 0
}

func (v Value) Equal(y Value) bool {
	return bytes.Equal([]byte(v), []byte(y))
}
