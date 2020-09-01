/**
 * Author: richen
 * Date: 2020-08-06 17:04:25
 * LastEditTime: 2020-08-07 14:24:13
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import "fmt"

func add(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return sum, sub
}

func main() {
	sum, sub := add(1, 2)
	fmt.Println(sum, sub)
}
