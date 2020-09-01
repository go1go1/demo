/**
 * Author: richen
 * Date: 2020-08-10 15:03:19
 * LastEditTime: 2020-08-10 17:58:07
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Addr string
	Post int
}

type User struct {
	Name string
	Age  int
	Addr string
	*Address
}

type User2 struct {
	Name string `json:"username",db:"user_name"`
	Age  int
	Addr string
	Address
}

func NewUser(name string, age int, addr string, post int) *User {
	user := &User{
		Name:    name,
		Age:     age,
		Address: &Address{},
	}
	user.Addr = addr //此处优先查找结构体本身的Addr字段，如果不存在继续向嵌套内寻找
	// Address为指针类型，需要初始化

	//第一种方式
	// user.Address = &Address{
	// 	Addr: "01",
	// 	Post: post,
	// }

	//第二种方式 前面必须初始化 Address: &Address{},
	user.Post = post

	return user
}

func testStruct() {
	var user User

	user.Addr = "dd"
	fmt.Printf("user:%#v\n", user)

	var u1 = &User{} //u1的类型是指针
	u1.Addr = "ee"   //golang语法糖，实际上是(*u1).Addr = "ee"
	fmt.Printf("u1:%#v\n", u1)

	var u2 = NewUser("aa", 18, "ff", 9)
	fmt.Printf("u2:%#v\n", u2.Post)

	var user2 User2
	user2.Post = 88 //User2未使用指针，可以直接赋值
	user2.Addr = "wwww"
	fmt.Printf("user2:%#v\n", user2)

	//struct tag
	data, _ := json.Marshal(user2)
	fmt.Printf("json string:%s\n", data)
}

func main() {
	testStruct()
}
