/**
 * Author: richen
 * Date: 2020-08-07 19:41:05
 * LastEditTime: 2020-08-07 20:06:36
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	lens int
	str  string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func parseArgs() {
	flag.IntVar(&lens, "l", 16, "-l 密码长度")
	flag.StringVar(&str, "t", "num", `
	-t 指定密码生产的字符
	 num: 只使用数字0-9
	 char: 只使用英文字母
	 mix: 使用数字和字母
	 advance: 字母以及特殊字符
	`)
	flag.Parse()
}

// 生成密码
func generatePasswd() string {
	var pass []byte = make([]byte, lens, lens)
	var sourceStr string
	if str == "num" {
		sourceStr = NumStr
	} else if str == "char" {
		sourceStr = CharStr
	} else if str == "mix" {
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr)
	} else if str == "advance" {
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr, SpecStr)
	} else {
		sourceStr = NumStr
	}

	// fmt.Printf("%s, %d\n", sourceStr, lens)

	for i := 0; i < lens; i++ {
		index := rand.Intn(len(sourceStr))
		// fmt.Printf("%d", index)
		pass[i] = sourceStr[index]
	}
	return string(pass)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	parseArgs()

	fmt.Printf("length: %d, charset: %s\n", lens, str)

	passwd := generatePasswd()

	fmt.Printf("NewPasswd: %s\n", passwd)
}
