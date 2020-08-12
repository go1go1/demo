/**
 * Author: richen
 * Date: 2020-08-10 11:48:25
 * LastEditTime: 2020-08-10 14:14:33
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import (
	"fmt"
	"strings"
)

//统计一个字符串每个单词出现的次数，比如："how do you do" => how=1 do=2 you=1
func ex1(str string) map[string]int {
	var result map[string]int = make(map[string]int, 128)
	words := strings.Split(str, " ")
	for _, v := range words {
		count, ok := result[v]
		if !ok {
			result[v] = 1
		} else {
			result[v] = count + 1
		}
	}
	return result
}

func studentStore() {

	var stuMap map[int]map[string]interface{}
	stuMap = make(map[int]map[string]interface{}, 16)
	// 学生id=1,姓名=stu01,分数=78.2,年龄=18
	var (
		id    = 1
		name  = "stu01"
		score = 78.2
		age   = 18
	)

	value, ok := stuMap[id]
	if !ok {
		value = make(map[string]interface{}, 8)
	}
	value["name"] = name
	value["id"] = id
	value["score"] = score
	value["age"] = age
	stuMap[id] = value

	fmt.Printf("stuMap: %#v\n", stuMap)

}

func main() {
	// fmt.Printf("result: %#v\n", ex1("how do you do do you like me"))

	studentStore()
}
