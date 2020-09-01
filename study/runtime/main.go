package main

import (
	"fmt"
	"runtime"
	"time"
)

func calc(i int) {
	for {
		i++
	}
}

func main() {
	cpu := runtime.NumCPU() //CPU核数
	fmt.Printf("cpu: %d\n", cpu)

	runtime.GOMAXPROCS(1) //设置最多使用1个核

	for i := 0; i < 10; i++ {
		go calc(i)
	}

	time.Sleep(time.Minute)
}
