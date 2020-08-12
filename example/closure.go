/**
 * Author: richen
 * Date: 2020-08-07 15:11:04
 * LastEditTime: 2020-08-07 15:28:09
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import (
	"fmt"
	"time"
)

func adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func adder2(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

// goroutine 使用闭包陷阱
func clos() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}

//闭包
func main() {
	// var f = adder2(1)

	// fmt.Println(f(1))
	// fmt.Println(f(20))
	// fmt.Println(f(30))

	clos()
}
