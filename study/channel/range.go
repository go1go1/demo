package main

import "fmt"

//channel的for range

func producer1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go producer1(ch)
	//使用for-range好处是不用判断管道是否被关闭
	for v := range ch {
		fmt.Printf("received: %v\n", v)
	}
}
