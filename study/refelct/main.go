/**
 * Author: richen
 * Date: 2020-08-12 09:43:21
 * LastEditTime: 2020-08-13 17:40:09
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import (
	"fmt"
	"reflect"
)

// 变量的内在机制
// A、类型信息，这部分是元信息，是预先定义好的
// B、值类型， 这部分是程序运行工程中，动态改变的

// 反射与空接口
// A、空接口可以存储任何类型的变量
// B、在运行时动态获取变量的类型和值信息，就叫反射

// reflect.Method 获取方法
// reflect.Name
// reflect.Kind
// reflect.ValueOf

func reflect_ex1(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	fmt.Printf("type of a:%v\n a=%v\n", t.Kind(), v)

	switch t.Kind() {
	case reflect.Int64:
		fmt.Printf("%d", v.Int())
	case reflect.Float64:
		fmt.Printf("%v", v.Float())
	case reflect.Ptr:
		v.Elem().SetFloat(6.8) //如果类型为指针，setFloat改变值的时候需要使用Elem来获取值
	}
}

type Student struct {
	Name string `json:"name" db:"name"`
	Sex  int
	Age  int
	xxx  string //结构体私有属性无法反射获取值，但是可以获取类型和名称
}

// 此处要注意：方法绑定在实体上时，通过实体反射，如果绑定在指针上(s *Student)，要通过指针来反射
func (s Student) Test() {
	fmt.Println("Student.Test(), reflect call")
}
func (s Student) test2(a string) {
	fmt.Printf("Student.test2(), %s\n", a)
}
func (s Student) Test3(a string) {
	fmt.Printf("Student.Test3(), %s\n", a)
}

// 反射结构体类型
func getStructType(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()
	k := t.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("s is %v\n", t)
	case reflect.Float64:
		fmt.Printf("s is %v\n", t)
	case reflect.Struct:
		fmt.Printf("s is %v\n", t)
		fmt.Printf("field num of s is %d\n", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			fmt.Printf("name: %s\n", t.Field(i).Name)
			fmt.Printf("type:%s\n", f.Type().Name())

			// fmt.Printf("value:%v\n", f.Interface()) //私有属性无法获取值
		}

		for y := 0; y < v.NumMethod(); y++ {
			m := t.Method(y)
			fmt.Printf("struct %d method:%s, type: %s\n", y, m.Name, m.Type)
		}
	}

	//通过反射进行调用方法
	//通过reflect.Value获取对应的方法并调用
	m := v.MethodByName("Test")
	var args []reflect.Value
	m.Call(args) //Test无参数,所以args是空切片

	//n := v.MethodByName("test2") //test2为私有方法，反射无法调用
	n := v.MethodByName("Test3")
	var args2 []reflect.Value
	var name string = "reflect call"
	args2 = append(args2, reflect.ValueOf(name)) //将参数放入切片，反射调用方法
	n.Call(args2)

}

//获取结构体的标签
func getStructFlag() {
	var s Student
	s.Name = "AAAA"

	v := reflect.ValueOf(s)
	t := v.Type()

	field0 := t.Field(0)

	fmt.Printf("tag json=%s\n", field0.Tag.Get("json"))
	fmt.Printf("tag db=%s\n", field0.Tag.Get("db"))

}

func main() {
	reflect_ex1(12)

	var s Student
	getStructType(s)

	getStructFlag()
}
