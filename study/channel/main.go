package main

import (
	"fmt"
	"time"
)

/**
channel本质上是一个队列，是一个容器

定义的时候，需要指定容器总元素的类型

var 变量名 chan 数据类型
*/

func hasCacheChannel() {
	var c chan int
	fmt.Printf("c=%v\n", c)
	c = make(chan int, 1) //缓冲区为1的队列
	fmt.Printf("c=%v\n", c)

	//入队操作
	c <- 100
	//出队操作
	data := <-c
	fmt.Printf("%v\n", data)
}

func product(c chan int) {
	for i := 1; i <= 10; i++ {
		c <- i + 1
	}
}

func consume(c chan int) {
	data := <-c
	fmt.Printf("%v\n", data)
}

func main() {
	var cc chan int = make(chan int, 5)
	go product(cc)
	go consume(cc)

	time.Sleep(time.Minute)
}
