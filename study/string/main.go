/**
 * Author: richen
 * Date: 2020-08-06 18:46:49
 * LastEditTime: 2020-08-07 14:45:08
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import "fmt"

func testString() {
	// 字符串底层就是byte数组，所以可以和[]byte类型互相转换
	var str = "hello"

	fmt.Printf("str[0] = %c\n", str[0])
	// string的长度就是[]byte的长度
	fmt.Printf("len(str)= %d\n", len(str))

	for index, val := range str {
		fmt.Printf("str[%d] = %c\n", index, val)
	}
	// 字符串一旦初始化不能修改
	// str[0] = "a" //error

	var byteSlice []byte
	byteSlice = []byte(str)
	byteSlice[0] = 'a' // byte用''单引号包括字符 不然报错cannot use "c" (type string) as type byte in assignment
	str = string(byteSlice)

	fmt.Println("after modify:", str)

	str = "少林之巅" //中文3个字节
	fmt.Printf("len(str)=%d\n", len(str))

	//rune类型用来表示utf8字符，一个rune字符由1个或多个byte组成
	runeSlice := []rune(str)
	fmt.Printf("str长度: %d；len(str)=%d\n", len(runeSlice), len(str))
}

// reverseString1 字符串逆序
func reverseString1() {
	var str = "hello"
	bytes := []byte(str)

	for i := 0; i < len(str)/2; i++ {
		tmp := bytes[len(str)-i-1]
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = tmp
	}
	str = string(bytes)
	fmt.Println(str)
}

// reverseString2 字符串逆序(含中文)
func reverseString2(str string) string {
	// var str = "hello中文中文"
	bytes := []rune(str)

	for i := 0; i < len(bytes)/2; i++ {
		tmp := bytes[len(bytes)-i-1]
		bytes[len(bytes)-i-1] = bytes[i]
		bytes[i] = tmp
	}
	str = string(bytes)
	fmt.Println(str)
	return str
}

// reverseString3 判断字符串是否逆序(含中文)
func reverseString3() {
	var str = "上海自来水来自海上"

	str2 := reverseString2(str)
	if str2 == str {
		fmt.Println("是回文")
	} else {
		fmt.Println("不是回文")
	}
}

// 计算字符串中字符、数字、空格、中文数量
func calc(str string) (charCount, numCount, spaceCount, otherCount int) {
	utfChars := []rune(str)

	for i := 0; i < len(utfChars); i++ {
		if (utfChars[i] >= 'a' && utfChars[i] <= 'z') || (utfChars[i] >= 'a' && utfChars[i] <= 'Z') {
			charCount++
			continue
		}

		if utfChars[i] > '0' && utfChars[i] <= '9' {
			numCount++
			continue
		}

		if utfChars[i] == ' ' {
			spaceCount++
			continue
		}
		otherCount++
	}
	return
}

func main() {
	// testString()
	// reverseString3()

	charCount, numCount, spaceCount, otherCount := calc("wofsldfj    时代峰峻类是否李经理 dsfs了 水电费9999")
	fmt.Printf("charCount: %d, numCount: %d, spaceCount: %d, otherCount: %d\n", charCount, numCount, spaceCount, otherCount)
}
