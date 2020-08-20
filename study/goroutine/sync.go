package main

import (
	"fmt"
	"time"
)

func hello(c chan bool) {
	fmt.Println("hello goroutine")

	time.Sleep(3 * time.Second)

	c <- true
}

func main() {
	var exChan chan bool
	exChan = make(chan bool)

	go hello(exChan)

	<-exChan //阻塞，等待channel内的消息
	fmt.Println("main thread terminate")

}
