/**
 * Author: richen
 * Date: 2020-08-10 18:22:56
 * LastEditTime: 2020-08-10 18:28:06
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import "fmt"

type Animal struct {
	Name string
	Sex  string
}

func (a *Animal) Talk() {
	fmt.Printf("talk : it's %s\n", a.Name)
}

type Dog struct {
	Feet string
	*Animal
}

func (d *Dog) Talk() {
	fmt.Printf("dog talk : it's %s\n", d.Name)
}

func test1() {
	var dog = &Dog{
		Feet: "4",
		Animal: &Animal{
			Name: "dog",
			Sex:  "xiong",
		},
	}

	dog.Talk()

}

func main() {
	test1()
}
