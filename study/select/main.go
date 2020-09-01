package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server2"
}

func test1() {
	out1 := make(chan string)
	out2 := make(chan string)

	go server1(out1)
	go server2(out2)

	s1 := <-out1
	fmt.Println(s1)
	//此处顺序执行，因为server1执行阻塞了6秒，虽然server2只有3秒阻塞，仍然会等待6秒
	s2 := <-out2
	fmt.Println(s2)
}

func server3(ch chan string) {
	time.Sleep(time.Second * 6)
	ch <- "response from server3"
}
func server4(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "response from server4"
}

/**
select :

1、同时监听一个或多个channel，直到其中一个channel ready
2、如果其中多个channel同时ready，随机选择一个进行操作
3、语法和switch case类似

*/
func test2() {
	out1 := make(chan string)
	out2 := make(chan string)

	go server3(out1)
	go server4(out2)

	//使用select避免不同goroutine互相阻塞
	select {
	case s1 := <-out1:
		fmt.Println(s1)
	case s2 := <-out2:
		fmt.Println(s2)
	}

}

/**

default分支，当所有分支channel都没有ready，则自动执行default分支
作用:
A、判断channel是否满了
B、判断channel是否为空

*/
func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write success")
		default:
			fmt.Println("channel is full")
		}

		time.Sleep(time.Millisecond * 500)
	}
}
func test3() {
	out1 := make(chan string, 10)

	go write(out1)

	for s := range out1 {
		fmt.Printf("recv:%s\n", s)
		time.Sleep(time.Second)
	}
}

func main() {
	//test1()
	//test2()
	test3()
}
