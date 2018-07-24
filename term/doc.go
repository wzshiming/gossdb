package term

var welcome = `SSDB (cli) - ssdb command line tool.
Copyright (c) 2018 github.com/wzshiming

	'help' for help, 'quit' to quit.
	
ssdb-server %s`

var usage = `- Server
  - auth password - Authenticate the connection.
  - dbsize - Return the approximate size of the database.
  - flushdb [type] - Delete all data in ssdb server.
  - info [opt] - Return the information of server.
  - slaveof id host port [auth last_seq last_key] - Start a replication slave.
- IP Filter
  - list_allow_ip ip_rule - List allow ip rules.
  - add_allow_ip ip_rule - add allow ip rules.
  - del_allow_ip ip_rule - del allow ip rules.
  - list_deny_ip ip_rule - list deny ip rules.
  - add_deny_ip ip_rule - add deny ip rules.
  - del_deny_ip ip_rule - del deny ip rules.
- Key Value
  - set key value - Set the value of the key.
  - setx key value ttl - Set the value of the key, with a time to live.
  - setnx key value - Set the string value in argument as value of the key only if the key doesn"t exist.
  - expire key ttl - Set the time left to live in seconds, only for keys of KV type.
  - ttl key - Returns the time left to live in seconds, only for keys of KV type.
  - get key - Get the value related to the specified key
  - getset key value - Sets a value and returns the previous entry at that key.
  - del key - Delete specified key.
  - incr key [num] - Increment the number stored at key by num.
  - exists key - Verify if the specified key exists.
  - getbit key offset - Return a single bit out of a string.
  - setbit key offset val - Changes a single bit of a string. The string is auto expanded.
  - bitcount key [start] [end] - Count the number of set bits (population counting) in part of a string.
  - countbit key start size - Count the number of set bits (population counting) in part of a string.
  - substr key start size - Return part of a string.
  - strlen key - Return the number of bytes of a string.
  - keys key_start key_end limit - List keys in range (key_start, key_end].
  - rkeys key_start key_end limit - List keys in range (key_start, key_end], in reverse order.
  - scan key_start key_end limit - List key-value pairs with keys in range (key_start, key_end].
  - rscan key_start key_end limit - List key-value pairs with keys in range (key_start, key_end], in reverse order.
  - multi_set key1 value1 key2 value2 ... - Set multiple key-value pairs(kvs) in one method call.
  - multi_get key1 key2 ... - Get the values related to the specified multiple keys
  - multi_del key1 key2 ... - Delete specified multiple keys.
- Hashmap
  - hset name key value - Set the string value in argument as value of the key of a hashmap.
  - hget name key - Get the value related to the specified key of a hashmap
  - hdel name key - Delete specified key in a hashmap.
  - hincr name key [num] - Increment the number stored at key in a hashmap by num
  - hexists name key - Verify if the specified key exists in a hashmap.
  - hsize name - Return the number of key-value pairs in the hashmap.
  - hlist name_start name_end limit - List hashmap names in range (name_start, name_end].
  - hrlist name_start name_end limit - List hashmap names in range (name_start, name_end].
  - hkeys name key_start key_end - List keys of a hashmap in range (key_start, key_end].
  - hgetall name - Returns the whole hash, as an array of strings indexed by strings.
  - hscan name key_start key_end limit - List key-value pairs of a hashmap with keys in range (key_start, key_end].
  - hrscan name key_start key_end limit - List key-value pairs with keys in range (key_start, key_end], in reverse order.
  - hclear name - Delete all keys in a hashmap.
  - multi_hset name key1 value1 key2 value2 ... - Set multiple key-value pairs(kvs) of a hashmap in one method call.
  - multi_hget name key1 key2 ... - Get the values related to the specified multiple keys of a hashmap.
  - multi_hdel name key1 key2 ... - Delete specified multiple keys in a hashmap.
- Sorted Set
  - zset name key score - Set the score of the key of a zset.
  - zget name key - Get the score related to the specified key of a zset
  - zdel name key - Delete specified key of a zset.
  - zincr name key num - Increment the number stored at key in a zset by num.
  - zexists name key - Verify if the specified key exists in a zset.
  - zsize name - Return the number of pairs of a zset.
  - zlist name_start name_end limit - List zset names in range (name_start, name_end].
  - zrlist name_start name_end limit - List zset names in range (name_start, name_end], in reverse order.
  - zkeys name key_start score_start score_end limit - List keys in a zset.
  - zscan name key_start score_start score_end limit - List key-score pairs where key-score in range (key_start+score_start, score_end].
  - zrscan name key_start score_start score_end limit - List key-score pairs of a zset, in reverse order. See method zkeys().
  - zrank name key - Returns the rank(index) of a given key in the specified sorted set.
  - zrrank name key - Returns the rank(index) of a given key in the specified sorted set, in reverse order.
  - zrange name offset limit - Returns a range of key-score pairs by index range [offset, offset + limit).
  - zrrange name offset limit - Returns a range of key-score pairs by index range [offset, offset + limit), in reverse order.
  - zclear name - Delete all keys in a zset.
  - zcount name score_start score_end - Returns the number of elements of the sorted set stored at the specified key which have scores in the range [score_start,score_end].
  - zsum name score_start score_end - Returns the sum of elements of the sorted set stored at the specified key which have scores in the range [score_start,score_end].
  - zavg name score_start score_end - Returns the average of elements of the sorted set stored at the specified key which have scores in the range [score_start,score_end].
  - zremrangebyrank name start end - Delete the elements of the zset which have rank in the range [start,end].
  - zremrangebyscore name start end - Delete the elements of the zset which have score in the range [start,end].
  - zpop_front name limit - Delete limit elements from front of the zset.
  - zpop_back name limit - Delete limit elements from back of the zset.
  - multi_zset name key1 score1 key2 score2 ... - Set multiple key-score pairs(kvs) of a zset in one method call.
  - multi_zget name key1 key2 ... - Get the values related to the specified multiple keys of a zset.
  - multi_zdel name key1 key2 ... - Delete specified multiple keys of a zset.
- List
  - qpush_front name item1 item2 ... - Adds one or more than one element to the head of the queue.
  - qpush_back name item1 item2 ... - Adds an or more than one element to the end of the queue.
  - qpop_front name size - Pop out one or more elements from the head of a queue.
  - qpop_back name size - Pop out one or more elements from the tail of a queue.
  - qpush name item1 item2 ... - Alias of qpush_back.
  - qpop name size - Alias of qpop_front.
  - qfront name - Returns the first element of a queue.
  - qback name - Returns the last element of a queue.
  - qsize name - Returns the number of items in the queue.
  - qclear name - Clear the queue.
  - qget name index - Returns the element a the specified index(position).
  - qset name index val - Description
  - qrange name offset limit - Returns a portion of elements from the queue at the specified range [offset, offset + limit].
  - qslice name begin end - Returns a portion of elements from the queue at the specified range [begin, end].
  - qtrim_front name size - Remove multi elements from the head of a queue.
  - qtrim_back name size - Remove multi elements from the tail of a queue.
  - qlist name_start name_end limit - List list/queue names in range (name_start, name_end].
  - qrlist name_start name_end limit - List list/queue names in range (name_start, name_end], in reverse order.

# display ssdb-server status
	info

see http://ssdb.io/docs/ for commands details

press 'quit' and Enter to quit.
`

var kind = map[string]string{
	"zsum":             "one",
	"zsize":            "one",
	"zset":             "none",
	"zscan":            "map",
	"zremrangebyscore": "none",
	"zremrangebyrank":  "none",
	"zrank":            "one",
	"zrange":           "map",
	"zrscan":           "map",
	"zrrank":           "one",
	"zrrange":          "map",
	"zrlist":           "list",
	"zpopfront":        "map",
	"zpopback":         "map",
	"zlist":            "list",
	"zkeys":            "list",
	"zincr":            "one",
	"zget":             "one",
	"zexists":          "one",
	"zdel":             "none",
	"zcount":           "one",
	"zclear":           "none",
	"zavg":             "one",
	"ttl":              "one",
	"substr":           "one",
	"strlen":           "one",
	"setx":             "none",
	"setnx":            "one",
	"setbit":           "one",
	"set":              "none",
	"scan":             "map",
	"rscan":            "map",
	"rkeys":            "list",
	"qtrimfront":       "one",
	"qtrimback":        "one",
	"qsize":            "one",
	"qset":             "none",
	"qrlist":           "list",
	"qpushfront":       "one",
	"qpushback":        "one",
	"qlist":            "list",
	"qget":             "one",
	"qfront":           "one",
	"qclear":           "none",
	"qback":            "one",
	"multizset":        "none",
	"multizget":        "map",
	"multizdel":        "none",
	"multiset":         "none",
	"multihset":        "none",
	"multihget":        "map",
	"multihdel":        "none",
	"multiget":         "map",
	"multidel":         "none",
	"keys":             "list",
	"info":             "info",
	"incr":             "one",
	"hsize":            "one",
	"hset":             "one",
	"hscan":            "map",
	"hrscan":           "map",
	"hrlist":           "list",
	"hlist":            "list",
	"hkeys":            "list",
	"hincr":            "one",
	"hgetall":          "map",
	"hget":             "one",
	"hexists":          "one",
	"hdel":             "one",
	"hclear":           "none",
	"getset":           "one",
	"getbit":           "one",
	"get":              "one",
	"expire":           "one",
	"exists":           "one",
	"del":              "none",
	"dbsize":           "one",
	"countbit":         "one",
	"bitcount":         "one",
	"auth":             "none",
}
