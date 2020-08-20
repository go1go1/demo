package main

import "fmt"

//channel的关闭

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //如果不关闭会导致消费死循环
}

func consumer(ch chan int) {
	for {
		v, ok := <-ch //使用common-ok表达式判断管道是否被关闭
		if ok == false {
			fmt.Println("chan is closed")
			break
		}
		fmt.Printf("received: %v\n", v)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)
	consumer(ch)
}
