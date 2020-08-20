package main

import (
	"fmt"
	"sync"
	"time"
)

/**
读写锁

读是并发的
读的时候不能写，可以多次读；写的时候不能读，其他人也不能写

*/
var (
	n      int
	wgg    sync.WaitGroup
	rwlock sync.RWMutex
)

func main() {
	wgg.Add(1)
	go write()

	for i := 0; i < 10; i++ {
		wgg.Add(1)
		fmt.Printf("goroutine: %d, n=%d\n", i, n)
		go read()
	}

	wgg.Wait()

}

func write() {
	rwlock.Lock()
	n = n + 1
	fmt.Printf("write-> n=%d\n", n)
	time.Sleep(time.Second * 10)
	rwlock.Unlock()
	wgg.Done()
}

func read() {
	rwlock.RLock()
	fmt.Printf("read-> x=%d\n", n)
	time.Sleep(time.Second)
	rwlock.RUnlock()
	wgg.Done()
}
