/**
 * Author: richen
 * Date: 2020-08-10 18:03:57
 * LastEditTime: 2020-08-10 18:18:16
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import "fmt"

type People struct {
	Name    string
	Country string
}

func (p *People) Print() {
	fmt.Printf("method print:%s,%s\n", p.Name, p.Country)
}

// 内置类型或者其他包定义的类型，不能直接扩展方法，但是可以使用别名进行扩展
type Integer int

func (i *Integer) Print() {
	fmt.Printf("i=%d\n", *i)
}

func main() {
	var p1 = &People{
		Name:    "01",
		Country: "cn",
	}

	p1.Print()

	var p2 = &People{
		Name:    "02",
		Country: "en",
	}

	p2.Print()
	p1.Print()

	fmt.Printf("p1:%p\n", p1)
	fmt.Printf("p2:%p\n", p2)

	var a Integer = 9
	a.Print()

}
