# SSDB client for golang

This is the SSDB client and command-line tool.  
Fork from [gossdb](https://github.com/ssdb/gossdb), but because it does not conform to the style of Golang and is not actively developed, the SSDB is rewrited.  

[![Go Report Card](https://goreportcard.com/badge/github.com/wzshiming/ssdb)](https://goreportcard.com/report/github.com/wzshiming/ssdb)
[![GoDoc](https://pkg.go.dev/badge/github.com/wzshiming/ssdb)](https://pkg.go.dev/github.com/wzshiming/ssdb)
[![GitHub license](https://img.shields.io/github/license/wzshiming/ssdb.svg)](https://github.com/wzshiming/ssdb/blob/master/LICENSE)

- [English](https://github.com/wzshiming/ssdb/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/ssdb/blob/master/README_cn.md)

## example

``` golang
package main

import (
	"fmt"

	"github.com/wzshiming/ssdb"
)

func main() {
	db, err := ssdb.Connect(
		ssdb.Addr("127.0.0.1:8888"),
		ssdb.Auth("password"),
		// or ssdb.URL("ssdb://127.0.0.1:8888?Auth=password"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Set("a", ssdb.Value("xxx"))
	if err != nil {
		fmt.Println(err)
		return
	}

	val, err := db.Get("a")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)

	err = db.Del("a")
	if err != nil {
		fmt.Println(err)
		return
	}

	val, err = db.Get("a")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)

	err = db.ZSet("z", "a", 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.MultiZSet("z", map[string]int64{
		"b": -1,
		"c": 5,
		"d": 3,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := db.ZRange("z", 0, 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range resp {
		fmt.Printf("  %v : %v\n", k, v)
	}

	return
}
```

Command-line tool

``` logs

go get -u -v github.com/wzshiming/ssdb/cmd/ssdb

ssdb -p password 127.0.0.1:8888

SSDB (cli) - ssdb command line tool.
Copyright (c) 2018 wzshiming

        'help' for help, 'quit' to quit.

ssdb-server 1.9.4
SSDB 127.0.0.1:8888 > set a xxx

1
1 result(s) (3ms)
SSDB 127.0.0.1:8888 > get a

xxx
1 result(s) (3ms)
SSDB 127.0.0.1:8888 > del a

1
1 result(s) (5ms)
SSDB 127.0.0.1:8888 > get a

not found
(4ms)
SSDB 127.0.0.1:8888 > zset z a 3

0
1 result(s) (3ms)
SSDB 127.0.0.1:8888 > multi_zset z b -1 c 5 d 3

0
1 result(s) (5ms)
SSDB 127.0.0.1:8888 > zrange z 0 10

key value
--- -----
b   -1
a   3
d   3
c   5
4 result(s) (28ms)

```

## API Support

[Official Documents](http://ssdb.io/docs/commands/index.html)  
[API Documents](https://godoc.org/github.com/wzshiming/ssdb)  

- [x] Server
  - [x] auth password - Authenticate the connection.
  - [x] dbsize - Return the approximate size of the database.
  - [ ] flushdb [type] - Delete all data in ssdb server.
  - [x] info [opt] - Return the information of server.
  - [ ] slaveof id host port [auth last_seq last_key] - Start a replication slave.
- [ ] IP Filter
- [x] Key Value
  - [x] set key value - Set the value of the key.
  - [x] setx key value ttl - Set the value of the key, with a time to live.
  - [x] setnx key value - Set the string value in argument as value of the key only if the key doesn"t exist.
  - [x] expire key ttl - Set the time left to live in seconds, only for keys of KV type.
  - [x] ttl key - Returns the time left to live in seconds, only for keys of KV type.
  - [x] get key - Get the value related to the specified key
  - [x] getset key value - Sets a value and returns the previous entry at that key.
  - [x] del key - Delete specified key.
  - [x] incr key [num] - Increment the number stored at key by num.
  - [x] exists key - Verify if the specified key exists.
  - [x] getbit key offset - Return a single bit out of a string.
  - [x] setbit key offset val - Changes a single bit of a string. The string is auto expanded.
  - [x] bitcount key [start] [end] - Count the number of set bits (population counting) in part of a string.
  - [x] countbit key start size - Count the number of set bits (population counting) in part of a string.
  - [x] substr key start size - Return part of a string.
  - [x] strlen key - Return the number of bytes of a string.
  - [x] keys key_start key_end limit - List keys in range (key_start, key_end].
  - [x] rkeys key_start key_end limit - List keys in range (key_start, key_end], in reverse order.
  - [x] scan key_start key_end limit - List key-value pairs with keys in range (key_start, key_end].
  - [x] rscan key_start key_end limit - List key-value pairs with keys in range (key_start, key_end], in reverse order.
  - [x] multi_set key1 value1 key2 value2 ... - Set multiple key-value pairs(kvs) in one method call.
  - [x] multi_get key1 key2 ... - Get the values related to the specified multiple keys
  - [x] multi_del key1 key2 ... - Delete specified multiple keys.
- [x] Hashmap
  - [x] hset name key value - Set the string value in argument as value of the key of a hashmap.
  - [x] hget name key - Get the value related to the specified key of a hashmap
  - [x] hdel name key - Delete specified key in a hashmap.
  - [x] hincr name key [num] - Increment the number stored at key in a hashmap by num
  - [x] hexists name key - Verify if the specified key exists in a hashmap.
  - [x] hsize name - Return the number of key-value pairs in the hashmap.
  - [x] hlist name_start name_end limit - List hashmap names in range (name_start, name_end].
  - [x] hrlist name_start name_end limit - List hashmap names in range (name_start, name_end].
  - [x] hkeys name key_start key_end - List keys of a hashmap in range (key_start, key_end].
  - [x] hgetall name - Returns the whole hash, as an array of strings indexed by strings.
  - [x] hscan name key_start key_end limit - List key-value pairs of a hashmap with keys in range (key_start, key_end].
  - [x] hrscan name key_start key_end limit - List key-value pairs with keys in range (key_start, key_end], in reverse order.
  - [x] hclear name - Delete all keys in a hashmap.
  - [x] multi_hset name key1 value1 key2 value2 ... - Set multiple key-value pairs(kvs) of a hashmap in one method call.
  - [x] multi_hget name key1 key2 ... - Get the values related to the specified multiple keys of a hashmap.
  - [x] multi_hdel name key1 key2 ... - Delete specified multiple keys in a hashmap.
- [x] Sorted Set
  - [x] zset name key score - Set the score of the key of a zset.
  - [x] zget name key - Get the score related to the specified key of a zset
  - [x] zdel name key - Delete specified key of a zset.
  - [x] zincr name key num - Increment the number stored at key in a zset by num.
  - [x] zexists name key - Verify if the specified key exists in a zset.
  - [x] zsize name - Return the number of pairs of a zset.
  - [x] zlist name_start name_end limit - List zset names in range (name_start, name_end].
  - [x] zrlist name_start name_end limit - List zset names in range (name_start, name_end], in reverse order.
  - [x] zkeys name key_start score_start score_end limit - List keys in a zset.
  - [x] zscan name key_start score_start score_end limit - List key-score pairs where key-score in range (key_start+score_start, score_end].
  - [x] zrscan name key_start score_start score_end limit - List key-score pairs of a zset, in reverse order. See method zkeys().
  - [x] zrank name key - Returns the rank(index) of a given key in the specified sorted set.
  - [x] zrrank name key - Returns the rank(index) of a given key in the specified sorted set, in reverse order.
  - [x] zrange name offset limit - Returns a range of key-score pairs by index range [offset, offset + limit).
  - [x] zrrange name offset limit - Returns a range of key-score pairs by index range [offset, offset + limit), in reverse order.
  - [x] zclear name - Delete all keys in a zset.
  - [x] zcount name score_start score_end - Returns the number of elements of the sorted set stored at the specified key which have scores in the range [score_start,score_end].
  - [x] zsum name score_start score_end - Returns the sum of elements of the sorted set stored at the specified key which have scores in the range [score_start,score_end].
  - [x] zavg name score_start score_end - Returns the average of elements of the sorted set stored at the specified key which have scores in the range [score_start,score_end].
  - [x] zremrangebyrank name start end - Delete the elements of the zset which have rank in the range [start,end].
  - [x] zremrangebyscore name start end - Delete the elements of the zset which have score in the range [start,end].
  - [x] zpop_front name limit - Delete limit elements from front of the zset.
  - [x] zpop_back name limit - Delete limit elements from back of the zset.
  - [x] multi_zset name key1 score1 key2 score2 ... - Set multiple key-score pairs(kvs) of a zset in one method call.
  - [x] multi_zget name key1 key2 ... - Get the values related to the specified multiple keys of a zset.
  - [x] multi_zdel name key1 key2 ... - Delete specified multiple keys of a zset.
- [x] List
  - [x] qpush_front name item1 item2 ... - Adds one or more than one element to the head of the queue.
  - [x] qpush_back name item1 item2 ... - Adds an or more than one element to the end of the queue.
  - [x] qpop_front name size - Pop out one or more elements from the head of a queue.
  - [x] qpop_back name size - Pop out one or more elements from the tail of a queue.
  - [ ] qpush name item1 item2 ... - Alias of `qpush_back`.
  - [ ] qpop name size - Alias of `qpop_front`.
  - [x] qfront name - Returns the first element of a queue.
  - [x] qback name - Returns the last element of a queue.
  - [x] qsize name - Returns the number of items in the queue.
  - [x] qclear name - Clear the queue.
  - [x] qget name index - Returns the element a the specified index(position).
  - [x] qset name index val - Description
  - [x] qrange name offset limit - Returns a portion of elements from the queue at the specified range [offset, offset + limit].
  - [x] qslice name begin end - Returns a portion of elements from the queue at the specified range [begin, end].
  - [x] qtrim_front name size - Remove multi elements from the head of a queue.
  - [x] qtrim_back name size - Remove multi elements from the tail of a queue.
  - [x] qlist name_start name_end limit - List list/queue names in range (name_start, name_end].
  - [x] qrlist name_start name_end limit - List list/queue names in range (name_start, name_end], in reverse order.

## License

Licensed under the MIT License. See [LICENSE](https://github.com/wzshiming/ssdb/blob/master/LICENSE) for the full license text.
