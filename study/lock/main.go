package main

import (
	"fmt"
	"sync"
	"time"
)

/**
读写锁在读多写少的场景，性能是互斥锁的十几倍
*/

var (
	n        int
	wgg, egg sync.WaitGroup
	rowlock  sync.RWMutex //读写锁
	mutex    sync.Mutex   //互斥锁
)

func rwWrite() {
	rowlock.Lock()
	//fmt.Println("write lock")
	n = n + 1
	fmt.Printf("write-> n=%d\n", n)
	//time.Sleep(time.Second * 10)
	//fmt.Println("write unlock")
	rowlock.Unlock()
	wgg.Done()
}

func rwRead(i int) {
	for i := 0; i < 1000; i++ {
		rowlock.RLock()
		//fmt.Println("read lock")
		fmt.Printf("read-> x=%d\n", i)
		//time.Sleep(time.Second)
		//fmt.Println("read unlock")
		rowlock.RUnlock()
	}

	wgg.Done()
}

func testRW() {
	wgg.Add(1)
	go rwWrite()
	//time.Sleep(time.Millisecond * 5)

	for i := 0; i < 100; i++ {
		wgg.Add(1)
		go rwRead(i)
	}

	wgg.Wait()
}

func mutexWrite() {
	mutex.Lock()
	//fmt.Println("write lock")
	n = n + 1
	fmt.Printf("write-> n=%d\n", n)
	//time.Sleep(time.Second * 10)
	//fmt.Println("write unlock")
	mutex.Unlock()
	egg.Done()
}

func mutexRead(i int) {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		//fmt.Println("read lock")
		fmt.Printf("read-> x=%d\n", i)
		//time.Sleep(time.Second)
		//fmt.Println("read unlock")
		mutex.Unlock()
	}

	egg.Done()
}

func testMutex() {
	egg.Add(1)
	go mutexWrite()
	//time.Sleep(time.Millisecond * 5)

	for i := 0; i < 100; i++ {
		egg.Add(1)
		go mutexRead(i)
	}

	egg.Wait()
}

func main() {
	//读写锁
	start := time.Now().UnixNano()
	testRW()
	end := time.Now().UnixNano()
	cost1 := (end - start) / 1000000

	//互斥锁
	start = time.Now().UnixNano()
	testMutex()
	end = time.Now().UnixNano()
	cost2 := (end - start) / 1000000

	fmt.Printf("testRW cost :%dms, testMutex cost :%dms", cost1, cost2)
}
