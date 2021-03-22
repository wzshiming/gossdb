package main

import (
	"fmt"

	"github.com/wzshiming/ssdb"
)

func main() {
	db, err := ssdb.Connect(
		ssdb.Addr("127.0.0.1:8888"),
		ssdb.Auth("password"),
		ssdb.IgnoreGetNotFoundError(true),
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
