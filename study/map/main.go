/**
 * Author: richen
 * Date: 2020-08-10 10:50:26
 * LastEditTime: 2020-08-10 11:43:20
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import (
	"fmt"
	"sort"
)

func map1(a map[string]int) {

	fmt.Printf("map a=%#v\n", a)

	//访问map内不存在的key，值为类型默认值
	result := a["d"]
	fmt.Printf("result:=%d\n", result)

	//判断map是否存在key
	key := "e"
	result, ok := a[key]
	if ok == false {
		fmt.Printf("key %s is not exist\n", key)
	}

	//map遍历
	for k, v := range a {
		fmt.Printf("key:%s, value:%d\n", k, v)
	}

	//map插入
	for i := 0; i < 16; i++ {
		key := fmt.Sprintf("st%d", i)
		a[key] = i
	}
	fmt.Printf("map a=%#v\n", a)

	//map删除
	delete(a, "st11")
	fmt.Printf("map a=%#v\n", a)

	fmt.Printf("map a length:%d\n", len(a))

}

func modify(a map[string]int) {
	a["modify"] = 1000
}

//map有序输出
func mapSort(a map[string]int) {
	var keys []string = make([]string, 0, 20)
	for k, _ := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Printf("keys: %s, values: %d\n", v, a[v])
	}
}

//map类型切片
func mapSlice() {
	var mapSli []map[string]int

	mapSli = make([]map[string]int, 5, 16)
	for k, v := range mapSli {
		fmt.Printf("mapSli[%d]=%v\n", k, v)
	}

	mapSli[0] = make(map[string]int)
	mapSli[0]["s1"] = 1
	mapSli[0]["s2"] = 1
	mapSli[0]["s3"] = 1

	fmt.Printf("mapSli: %#v", mapSli)
}

func main() {
	var a map[string]int

	a = make(map[string]int, 16)
	fmt.Printf("map a=%#v\n", a)
	a["a"] = 1
	a["b"] = 2
	a["c"] = 3

	map1(a)

	//map是引用类型
	b := a
	modify(b)
	fmt.Printf("map b=%#v\n", b)
	fmt.Printf("map a=%#v\n", a)

	mapSort(a)
	mapSlice()
}
