package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**

原子操作

原子操作仅对基本类型有效

原子操作性能要比加锁好很多
*/

var (
	x     int32
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func add() {
	for i := 0; i < 5000; i++ {
		//x += 1
		//修改为原子操作
		atomic.AddInt32(&x, 1)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()

	wg.Wait()

	fmt.Printf("x = %d", x)
}
