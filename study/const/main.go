/**
 * Author: richen
 * Date: 2020-08-06 17:13:14
 * LastEditTime: 2020-08-06 17:17:59
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import "fmt"

const (
	A = iota // 0
	B        // 1
	C        // 2
	D = 8
	E        //未定义，默认等于上一行 8
	F = iota // iota 5

	G // iota 6
)

const (
	A1 = iota
	A2
)

func main() {
	fmt.Println("A: ", A)
	fmt.Println("B: ", B)
	fmt.Println("C: ", C)
	fmt.Println("D: ", D)
	fmt.Println("E: ", E)
	fmt.Println("F: ", F)
	fmt.Println("G: ", G)
}
