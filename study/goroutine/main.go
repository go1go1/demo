/**
 * Author: richen
 * Date: 2020-08-06 16:48:40
 * LastEditTime: 2020-08-06 17:01:22
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

/**
并发和并行

并发：同一时间段内执行多个操作
并行：同一时刻执行多个操作

线程是由操作系统进行管理，也就是处于内核态
线程之间进行切换，需要发生用户态到内核态的切换
当系统中运行大量线程，系统会变得非常慢
用户态的线程，由程序自行管理（调度和切换）而无需发生用户态和内核态切换，支持大量被创建，也叫协程。


抽象模型:

操作系统线程: M
用户态线程(goroutine): G
上下文对象：P

*/

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {

	go numbers()
	go alphabets()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\nmain exited")
}
