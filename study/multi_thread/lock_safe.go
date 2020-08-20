package main

import (
	"fmt"
	"sync"
)

/**
线程安全

 - 多个goroutine同时操作一个资源，这个资源又叫临界区
   十字路口，通过红绿灯实现线程安全； 火车上的厕所，通过互斥锁实现线程安全

*/
var (
	x     int
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func add() {
	for i := 0; i < 5000; i++ {
		mutex.Lock() //加锁
		x += 1
		mutex.Unlock() //释放锁
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
