/**
 * Author: richen
 * Date: 2020-08-07 11:41:43
 * LastEditTime: 2020-08-07 14:32:40
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import (
	"fmt"
)

// 九九乘法表
func nine() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, j*i)
		}
		fmt.Println()
	}
}

func isShux(n int) bool {
	first := n % 10
	second := (n / 10) % 10
	third := (n / 100) % 10

	sum := first*first*first + second*second*second + third*third*third
	if sum == n {
		return true
	}
	return false
}

// 水仙花数，所谓水仙花数是一个三位数，其各位数字的立方和等于该数本身，例如 153 = 1^3 + 5^3 + 3^3
func snum() {
	//100到1000之间
	for i := 100; i < 1000; i++ {
		if isShux(i) == true {
			fmt.Printf("数字[%d]是水仙花数\n", i)
		}
	}
}

func main() {
	// nine()
	snum()
}
