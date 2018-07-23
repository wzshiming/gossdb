# SSDB 的 golang 客户端

从官方客户端派生更符合golang风格并支持连接池。

[![Go Report Card](https://goreportcard.com/badge/github.com/wzshiming/ssdb)](https://goreportcard.com/report/github.com/wzshiming/ssdb)
[![GoDoc](https://godoc.org/github.com/wzshiming/ssdb?status.svg)](https://godoc.org/github.com/wzshiming/ssdb)
[![GitHub license](https://img.shields.io/github/license/wzshiming/ssdb.svg)](https://github.com/wzshiming/ssdb/blob/master/LICENSE)

- [English](https://github.com/wzshiming/ssdb/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/ssdb/blob/master/README_cn.md)

## 示例

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

	err = db.Set("a", "xxx")
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

## API 支持

[官方文档](http://ssdb.io/docs/zh_cn/commands/index.html)  
[API 文档](https://godoc.org/github.com/wzshiming/ssdb)  

- [x] Server
  - [x] auth password - 向服务器校验访问密码.
  - [x] dbsize - 返回数据库占用空间的近似值, 以字节为单位.
  - [ ] flushdb [type] - 删除 SSDB 服务器的所有数据.
  - [x] info [opt] - 返回服务器的信息.
  - [ ] slaveof id host port [auth last_seq last_key] - 启动一个从库同步进程.
- [ ] IP Filter
- [x] Key Value
  - [x] set key value - 设置指定 key 的值内容.
  - [x] setx key value ttl - 设置指定 key 的值内容, 同时设置存活时间.
  - [x] setnx key value - 当 key 不存在时, 设置指定 key 的值内容. 如果已存在, 则不设置.
  - [x] expire key ttl - 设置 key(只针对 KV 类型) 的存活时间.
  - [x] ttl key - 返回 key(只针对 KV 类型) 的存活时间.
  - [x] get key - 获取指定 key 的值内容.
  - [x] getset key value - 更新 key 对应的 value, 并返回更新前的旧的 value.
  - [x] del key - 删除指定的 key.
  - [x] incr key [num] - 使 key 对应的值增加 num.
  - [x] exists key - 判断指定的 key 是否存在.
  - [x] getbit key offset - 获取字符串内指定位置的位值(BIT).
  - [x] setbit key offset val - 设置字符串内指定位置的位值(BIT), 字符串的长度会自动扩展.
  - [x] bitcount key [start] [end] - 计算字符串的子串所包含的位值为 1 的个数.
  - [x] countbit key start size - 计算字符串的子串所包含的位值为 1 的个数.
  - [x] substr key start size - 获取字符串的子串.
  - [x] strlen key - 计算字符串的长度(字节数).
  - [x] keys key_start key_end limit - 列出处于区间 (key_start, key_end] 的 key 列表.
  - [x] rkeys key_start key_end limit - 列出处于区间 (key_start, key_end] 的 key 列表, 反向.
  - [x] scan key_start key_end limit - 列出处于区间 (key_start, key_end] 的 key-value 列表.
  - [x] rscan key_start key_end limit - 列出处于区间 (key_start, key_end] 的 key-value 列表, 反向.
  - [x] multi_set key1 value1 key2 value2 ... - 批量设置一批 key-value.
  - [x] multi_get key1 key2 ... - 批量获取一批 key 对应的值内容.
  - [x] multi_del key1 key2 ... - 批量删除一批 key 和其对应的值内容.
- [x] Hashmap
  - [x] hset name key value - 设置 hashmap 中指定 key 对应的值内容.
  - [x] hget name key - 获取 hashmap 中指定 key 的值内容.
  - [x] hdel name key - 删除 hashmap 中的指定 key(删除整个 hashmap 用 hclear).
  - [x] hincr name key [num] - 使 hashmap 中的 key 对应的值增加 num.
  - [x] hexists name key - 判断指定的 key 是否存在于 hashmap 中.
  - [x] hsize name - 返回 hashmap 中的元素个数.
  - [x] hlist name_start name_end limit - 列出名字处于区间 (name_start, name_end] 的 hashmap.
  - [x] hrlist name_start name_end limit - 像 hrlist, 逆序.
  - [x] hkeys name key_start key_end - 列出 hashmap 中处于区间 (key_start, key_end] 的 key 列表.
  - [x] hgetall name - 返回整个 hashmap.
  - [x] hscan name key_start key_end limit - 列出 hashmap 中处于区间 (key_start, key_end] 的 key-value 列表.
  - [x] hrscan name key_start key_end limit - 像 hscan, 逆序.
  - [x] hclear name - 删除 hashmap 中的所有 key.
  - [x] multi_hset name key1 value1 key2 value2 ... - 批量设置 hashmap 中的 key-value.
  - [x] multi_hget name key1 key2 ... - 批量获取 hashmap 中多个 key 对应的权重值.
  - [x] multi_hdel name key1 key2 ... - 指删除 hashmap 中的 key.
- [x] Sorted Set
  - [x] zset name key score - 设置 zset 中指定 key 对应的权重值.
  - [x] zget name key - 获取 zset 中指定 key 的权重值.
  - [x] zdel name key - 获取 zset 中的指定 key.
  - [x] zincr name key num - 使 zset 中的 key 对应的值增加 num. 参数 num 可以为负数. 如果原来的值不是整数(字符串形式的整数), 它会被先转换成整数.
  - [x] zexists name key - 判断指定的 key 是否存在于 zset 中.
  - [x] zsize name - 返回 zset 中的元素个数.
  - [x] zlist - 列出名字处于区间 (name_start, name_end] 的 zset.
  - [x] zrlist - 像 zlist, 逆序.
  - [x] zkeys name key_start score_start score_end limit - 列出 zset 中的 key 列表.
  - [x] zscan name key_start score_start score_end limit - 列出 zset 中处于区间 (key_start+score_start, score_end] 的 key-score 列表.
  - [x] zrscan name key_start score_start score_end limit - 像 zscan, 逆序.
  - [x] zrank name key - 返回指定 key 在 zset 中的排序位置(排名), 排名从 0 开始.
  - [x] zrrank name key - 像 zrank, 逆序.
  - [x] zrange name offset limit - 根据下标索引区间 [offset, offset + limit) 获取 key-score 对, 下标从 0 开始.
  - [x] zrrange name offset limit - 像 zrange, 逆序.
  - [x] zclear name - 删除 zset 中的所有 key.
  - [x] zcount name start end - 返回处于区间 [start,end] key 数量.
  - [x] zsum name start end - 返回 key 处于区间 [start,end] 的 score 的和.
  - [x] zavg name start end - 返回 key 处于区间 [start,end] 的 score 的平均值.
  - [x] zremrangebyrank name start end - 删除位置处于区间 [start,end] 的元素.
  - [x] zremrangebyscore name start end - 删除权重处于区间 [start,end] 的元素.
  - [x] zpop_front name limit - 从 zset 首部删除 limit 个元素.
  - [x] zpop_back name limit - 从 zset 尾部删除 limit 个元素.
  - [x] multi_zset name key1 score1 key2 score2 ... - 批量设置 zset 中的 key-score.
  - [x] multi_zget name key1 key2 ... - 批量获取 zset 中多个 key 对应的权重值.
  - [x] multi_zdel name key1 key2 ... - 批量删除 zset 中的 key.
- [x] List
  - [x] qpush_front name item1 item2 ... - 往队列的首部添加一个或者多个元素.
  - [x] qpush_back name item1 item2 ... - 往队列的尾部添加一个或者多个元素.
  - [x] qpop_front name size - 从队列首部弹出最后一个或者多个元素.
  - [x] qpop_back name size - 从队列尾部弹出最后一个或者多个元素.
  - [ ] qpush name item1 item2 ... - 是 `qpush_back` 的别名..
  - [ ] qpop name size - 是 `qpop_front` 的别名..
  - [x] qfront name - 返回队列的第一个元素.
  - [x] qback name - 返回队列的最后一个元素.
  - [x] qsize name - 返回队列的长度.
  - [x] qclear name - 清空一个队列.
  - [x] qget name index - 返回指定位置的元素.
  - [x] qset name index val - 更新位于 index 位置的元素.
  - [x] qrange name offset limit - 返回下标处于区域 [offset, offset + limit] 的元素.
  - [x] qslice name begin end - 返回下标处于区域 [begin, end] 的元素. begin 和 end 可以是负数
  - [x] qtrim_front name size - 从队列头部删除多个元素.
  - [x] qtrim_back name size - 从队列头部删除多个元素.
  - [x] qlist name_start name_end limit - 列出名字处于区间 (name_start, name_end] 的 queue/list.
  - [x] qrlist name_start name_end limit - 像 qlist, 逆序.

## 许可证

软包根据MIT License。有关完整的许可证文本，请参阅[LICENSE](https://github.com/wzshiming/ssdb/blob/master/LICENSE)。
