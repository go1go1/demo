/**
 * Author: richen
 * Date: 2020-08-07 10:13:41
 * LastEditTime: 2020-08-07 11:42:05
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import (
	"fmt"
	"time"
)

func testTimse() {
	now := time.Now()
	fmt.Printf("current time: %+v \n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	// %02d不够两位补全
	fmt.Printf("current daytime: %02d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second)

	timestamp := now.Unix()
	fmt.Printf("timestamp: %d \n", timestamp)
}

// 时间戳转时间字符串
func getDayTimeFromTimestamp(t int64) {
	//第一个参数时间戳，第二个参数为纳秒
	ot := time.Unix(t, 0)
	year := ot.Year()
	month := ot.Month()
	day := ot.Day()
	hour := ot.Hour()
	minute := ot.Minute()
	second := ot.Second()

	aa := fmt.Sprintf("current daytime: %02d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second)

	fmt.Printf("%s", aa)
}

func processTask() {
	fmt.Println("do task")
}

// 定时器
func testTicker() {
	// 按照2秒
	ticker := time.Tick(2 * time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		processTask()
	}
}

//格式化ç
func testFormat() {
	now := time.Now()
	fmt.Println(now.Format("02/1/2006 3:04:05"))
	fmt.Println(now.Format("02/1/2006 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}

//计算程序执行耗时
func testCost() {
	start := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	cost := (end - start) / 1000
	fmt.Printf("code cost: %dus \n", cost)
}

func main() {
	// testTimse()

	// getDayTimeFromTimestamp(time.Now().Unix())
	// testTicker()

	// time.Duration用来表示纳秒
	// 一些常量
	// const (
	// 	Nanosecond  Duration = 1                  //纳秒
	// 	Microsecond          = 1000 * Nanosecond  //微妙
	// 	Millisecond          = 1000 * Microsecond //毫秒
	// 	Second               = 1000 * Millisecond //秒
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )
	// testFormat()

	testCost()

}
