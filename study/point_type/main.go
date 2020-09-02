/**
 * Author: richen
 * Date: 2020-08-10 09:59:05
 * LastEditTime: 2020-08-10 10:49:41
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import "fmt"

func swap(a *int, b *int) {
	fmt.Printf("a = %d, b= %d\n", a, b)
	*a, *b = *b, *a
	fmt.Printf("a = %d, b= %d\n", *a, *b)
}

func main() {
	// var a int32
	// a = 1000
	// fmt.Printf("addr of a:%p, a:%d\n", &a, a)

	// var b *int32
	// b = &a
	// fmt.Printf("addr of b:%p, b:%v\n", &b, b)

	// var b = 10
	// p := &b1
	// *p = 100
	// // modify(p)
	// fmt.Printf("b1: %v, p:%v\n", b1, *p)

	var a = 10
	var b = 20

	swap(&a, &b)
	fmt.Printf("a = %d, b= %d\n", a, b)
}

func modify(a *int) {
	*a = 100
}
