package main

import (
	"fmt"
	"github.com/gobestsdk/memredis"
	"time"
)

func main() {

	memredis.SET("l_log", "app.log", -1)
	memredis.SET("l_db", "app.s", time.Second*7)
	memredis.SET("l_dbf2", "app.s", time.Second*12)
	memredis.SET("xa", "app.f", time.Second*3)

	for {
		select {
		case <-time.After(time.Second * 1):
			fmt.Println("keys", memredis.Keys("l_*"))
		}
	}

}
