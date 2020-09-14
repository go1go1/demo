package main

import (
	"fmt"
	"time"
)

//生产者
func produceSushu(c chan int) {
	var i int = 1
	for {
		i = i + 1
		result := isPrime(i)
		if result {
			c <- i
		}
		time.Sleep(time.Second)
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//消费者
func consumeSushu(c chan int) {
	for v := range c {
		fmt.Printf("%d is prime\n", v)
	}
}

// dlv debug ./memory_session.go   开始调试
// b memory_session.go:16          在main.go文件16行断点
// b produceSushu        在produceSushu函数断点
// c                     执行到断点
// n                     单步跳过
// s                     单步进入
// p i                   打印i的值
// bt                    打印堆栈
// goroutines            查看goroutine列表
// goroutine 18          切换到goroutine18
// q                     退出调试

func main() {
	var intChan chan int = make(chan int, 1000)
	go produceSushu(intChan)
	go consumeSushu(intChan)

	time.Sleep(time.Hour)
}
