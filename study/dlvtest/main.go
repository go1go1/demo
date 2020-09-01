package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		var i int
		time.Sleep(5 * time.Second)
		i++

		curTime := time.Now()
		fmt.Printf("run %d count, cur time:%v\n", i, curTime)
	}
}
