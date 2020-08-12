/**
 * Author: richen
 * Date: 2020-08-07 18:57:17
 * LastEditTime: 2020-08-07 19:25:32
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import "fmt"

func test() {
	//基于数组初始化切片
	a := [3]int{1, 2, 3}
	b := a[1:2] // 数组start ~ end-1 (2-1)
	fmt.Printf("%v\n", b)

	b1 := a[1:] // 数组start ~ 结束
	fmt.Printf("%v\n", b1)

	b2 := a[:] // 数组
	fmt.Printf("%v\n", b2)

	c := []int{6, 7, 8}
	fmt.Printf("%v\n", c)

	//make创建切片
	var a1 []int

	a1 = make([]int, 5, 10) // 5 切片长度(创建时长度), 10 切片容量(底层数组容量)
	fmt.Printf("%v\n", a1)
}

func editSlice() {
	a := []int{1, 2, 3}
	for i := range a {
		a[i]++
	}

	fmt.Printf("%d\n", a)
}

// 切片扩容，扩容的策略是翻倍扩容
func testMake1() {
	var a []int
	a = make([]int, 1, 5)

	a[0] = 5
	// a[1] = 10 //error 切片长度为1，此处越界
	fmt.Printf("a = %v, addr:%p, len:%d, cap:%d\n", a, a, len(a), cap(a))
	for i := 0; i < 4; i++ {
		a = append(a, i+1)
		fmt.Printf("a = %v, addr:%p, len:%d, cap:%d\n", a, a, len(a), cap(a))
	}

	a = append(a, 14)
	fmt.Println("内存地址变化,发生扩容,容量翻倍===>")
	fmt.Printf("a = %v, addr:%p, len:%d, cap:%d\n", a, a, len(a), cap(a))

}

func main() {
	testMake1()
}
