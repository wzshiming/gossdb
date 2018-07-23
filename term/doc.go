package term

var usage = `
# display ssdb-server status
	info

see http://ssdb.io/docs/ for commands details

press 'quit' and Enter to quit.`

var welcome = `SSDB (cli) - ssdb command line tool.
Copyright (c) 2018 github.com/wzshiming

	'help' for help, 'quit' to quit.
	
ssdb-server %s`

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
