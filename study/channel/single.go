package main

import "fmt"

// 单向channel，读取权限控制

//只写 c chan<- int
func sendData(c chan<- int) {
	c <- 10
}

//只读 c <-chan int
func readData(c <-chan int) {
	fmt.Printf("%v", <-c)
}

func main() {
	chn := make(chan int, 1)
	go sendData(chn)
	readData(chn)
}
