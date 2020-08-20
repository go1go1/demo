package main

import (
	"fmt"
	"sync"
	"time"
)

/**
如何等待一组goroutine结束？
*/
//方法一，使用不带缓冲区的channel实现
func process(i int, ch chan bool) {
	fmt.Printf("started goroutine %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("goroutine %d ended\n", i)
	ch <- true
}

func go1() {
	goroutineNum := 3
	exChan := make(chan bool, goroutineNum)

	for i := 0; i < goroutineNum; i++ {
		go process(i, exChan)
	}

	for i := 0; i < goroutineNum; i++ {
		<-exChan
	}

	fmt.Printf("all goroutines finished.")
}

//方法二， 使用sync.WaitGroup实现, 推荐写法
func process2(i int, wg *sync.WaitGroup) {
	fmt.Printf("started goroutine %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("goroutine %d ended\n", i)
	wg.Done()
}
func go2() {
	goroutineNum := 3
	var wg sync.WaitGroup

	for i := 0; i < goroutineNum; i++ {
		wg.Add(i)
		go process2(i, &wg)
	}

	wg.Wait()
	fmt.Printf("all goroutines finished.")
}

func main() {
	//go1()
	go2()
}
