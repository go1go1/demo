/**
 * Author: richen
 * Date: 2020-08-10 18:30:45
 * LastEditTime: 2020-08-11 09:38:07
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	Name     string
	Count    int
	Students []*Student
}

type Student struct {
	Id   int
	Name string
	Sex  string
}

func testClass() {
	c := &Class{
		Name:  "101",
		Count: 0,
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name: fmt.Sprintf("stu%d", i),
			Sex:  "man",
			Id:   i,
		}
		c.Students = append(c.Students, stu)
	}
	//json序列化
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marchal failed")
		return
	}
	fmt.Printf("json: %s\n", string(data))

	//json反序列化
	c1 := &Class{}
	err = json.Unmarshal(data, c1)
	if err != nil {
		fmt.Println("json unmarchal failed")
		return
	}
	fmt.Printf("c1: %#v\n", c1)

	for _, v := range c1.Students {
		fmt.Printf("student: %#v\n", v)
	}
}

func main() {
	testClass()
}
