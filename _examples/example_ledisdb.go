package main

import (
	"fmt"

	"../format/as"
	"github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

func main() {
	key := []byte("Test")
	value := []byte("TESTING")
	l, _ := ledis.Open(config.NewConfigDefault())
	db, _ := l.Select(0)

	db.Set(key, value)

	fmt.Println(as.ToString(db.Get(key)))
}
