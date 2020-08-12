/**
 * Author: richen
 * Date: 2020-08-07 16:06:51
 * LastEditTime: 2020-08-07 18:55:03
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

//golang内基本数据类型都是值类型
import "fmt"

// 找出数组中为给定值的两个元素的下标，比如数组[1,3,5,8,7]，找出两个元素之和等于8的下标
func findArr(arr [5]int, target int) {
	// arr := [5]int{1, 3, 5, 8, 7}
	for i := 0; i < len(arr); i++ {
		temp := target - arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == temp {
				fmt.Printf("%d, %d\n", i, j)
			}
		}
	}
}

func main() {
	a := [3]int{10, 20, 30}
	b := a //数组为值类型，赋值时会进行拷贝

	b[0] = 100

	// fmt.Println(a)
	// fmt.Println(b)

	findArr([...]int{1, 3, 5, 8, 7}, 8)
}
