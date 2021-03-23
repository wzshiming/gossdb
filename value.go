package ssdb

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Values Value Slice
type Values []Value

func NewValues(arg []interface{}) (Values, error) {
	vs := make(Values, 0, len(arg))
	for _, v := range arg {
		d, err := NewValue(v)
		if err != nil {
			return nil, err
		}
		vs = append(vs, d)
	}
	return vs, nil
}

// String return first value
func (v Values) First() Value {
	if len(v) == 0 {
		return nil
	}
	return v[0]
}

// String return string
func (v Values) String() string {
	if len(v) == 0 {
		return ""
	}
	return strings.Join(v.Strings(), "\n")
}

// Strings get []string
func (v Values) Strings() []string {
	if len(v) == 0 {
		return nil
	}
	s := make([]string, 0, len(v))
	for _, v := range v {
		s = append(s, v.String())
	}
	return s
}

// MapStringInt get map[string]int64
func (v Values) MapStringInt() map[string]int64 {
	if len(v) == 0 {
		return nil
	}
	val := map[string]int64{}
	size := len(v)
	for i := 0; i+1 < size; i += 2 {
		val[v[i].String()] = v[i+1].Int()
	}
	return val
}

// MapStringValue get map[string]Value
func (v Values) MapStringValue() map[string]Value {
	if len(v) == 0 {
		return nil
	}
	val := map[string]Value{}
	size := len(v)
	for i := 0; i+1 < size; i += 2 {
		val[v[i].String()] = v[i+1]
	}
	return val
}

// Pairs get original key-value pairs
func (v Values) Pairs() Pairs {
	if len(v) == 0 {
		return nil
	}
	val := make(Pairs, 0, len(v)/2)
	size := len(v)
	for i := 0; i+1 < size; i += 2 {
		val = append(val, Pair{
			Key:   v[i],
			Value: v[i+1],
		})
	}
	return val
}

// Value return val
type Value []byte

func NewValue(arg interface{}) (Value, error) {
	switch arg := arg.(type) {
	case Value:
		return arg, nil
	case time.Duration:
		return Value(strconv.AppendUint(nil, uint64(arg/time.Second), 10)), nil
	case fmt.Stringer:
		return Value(arg.String()), nil
	case string:
		return Value(arg), nil
	case []byte:
		return Value(arg), nil
	case int:
		return Value(strconv.AppendInt(nil, int64(arg), 10)), nil
	case int8:
		return Value(strconv.AppendInt(nil, int64(arg), 10)), nil
	case int16:
		return Value(strconv.AppendInt(nil, int64(arg), 10)), nil
	case int32:
		return Value(strconv.AppendInt(nil, int64(arg), 10)), nil
	case int64:
		return Value(strconv.AppendInt(nil, int64(arg), 10)), nil
	case uint:
		return Value(strconv.AppendUint(nil, uint64(arg), 10)), nil
	case uint8:
		return Value(strconv.AppendUint(nil, uint64(arg), 10)), nil
	case uint16:
		return Value(strconv.AppendUint(nil, uint64(arg), 10)), nil
	case uint32:
		return Value(strconv.AppendUint(nil, uint64(arg), 10)), nil
	case uint64:
		return Value(strconv.AppendUint(nil, uint64(arg), 10)), nil
	case float32:
		return Value(strconv.AppendFloat(nil, float64(arg), 'f', -1, 64)), nil
	case float64:
		return Value(strconv.AppendFloat(nil, float64(arg), 'f', -1, 64)), nil
	case bool:
		if arg {
			return one, nil
		} else {
			return zero, nil
		}
	case nil:
		return Value(""), nil
	default:
		return nil, fmt.Errorf("error type")
	}
}

// String
func (v Value) String() string {
	if len(v) == 0 {
		return ""
	}
	return string(v)
}

// Duration get time.Duration
func (v Value) Duration() time.Duration {
	return time.Duration(v.Int()) * time.Second
}

// Int get int
func (v Value) Int() int64 {
	i, _ := strconv.ParseInt(string(v), 0, 0)
	return i
}

// Uint get uint
func (v Value) Uint() uint64 {
	i, _ := strconv.ParseUint(string(v), 0, 0)
	return i
}

// Float get float
func (v Value) Float() float64 {
	i, _ := strconv.ParseFloat(string(v), 0)
	return i
}

// Bool get bool
func (v Value) Bool() bool {
	i, _ := strconv.ParseBool(string(v))
	return i
}

// Bytes get bytes
func (v Value) Bytes() []byte {
	return []byte(v)
}

// IsEmpty is empty
func (v Value) IsEmpty() bool {
	return len(v) == 0
}

// Equal value equal
func (v Value) Equal(y Value) bool {
	return bytes.Equal([]byte(v), []byte(y))
}

// Pair the original key-value pair
type Pair struct {
	Key   Value
	Value Value
}

// Pairs is the list of pair
type Pairs []Pair
