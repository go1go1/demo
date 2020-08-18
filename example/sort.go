/**
 * Author: richen
 * Date: 2020-08-07 15:29:14
 * LastEditTime: 2020-08-07 19:39:02
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import (
	"fmt"
	"sort"
)

//插入排序
func insertSort(a [8]int) [8]int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				fmt.Println(i, ":", a)
			} else {
				break
			}
		}
	}
	return a
}

//选择排序
func selectSort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
				fmt.Println(i, ":", a)
			}
		}
	}
	return a
}

//冒泡排序
func bubbleSort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-(i+1); j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				fmt.Println(i, ":", a)
			}
		}
	}
	return a
}

// golang标准库sort排序
func standSort(a [8]int) {
	//var a [5]int = [5]int{5, 3, 4, 2, 1}
	sort.Ints(a[:])

	fmt.Println("a:", a)
}

func main() {
	var i = [8]int{8, 3, 2, 9, 4, 6, 10, 0}
	// fmt.Println(i)
	// fmt.Println(insertSort(i))
	// fmt.Println(selectSort(i))
	// fmt.Println(bubbleSort(i))
	standSort(i)
}
