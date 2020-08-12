/**
 * Author: richen
 * Date: 2020-08-11 11:11:38
 * LastEditTime: 2020-08-11 20:24:43
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package logi

import (
	"fmt"
	"os"
	"runtime"
)

type LogData struct {
	Message  string
	TimeStr  string
	ColorStr string
	LevelStr string
	FileName string
	FuncName string
	ResetStr string
	LineNo   int
}

// exists 判断所给路径文件/文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// mkDir 创建目录
func mkDir(path string) {
	// 创建文件夹
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Printf("mkdir failed![%v]\n", err)
	} else {
		fmt.Printf("mkdir success!\n")
	}
}

// fSize 获取文件大小
func fSize(f *os.File) int64 {
	statInfo, err := f.Stat()
	if err != nil {
		return 0
	}
	return statInfo.Size()
}

// getLineInfo
func getLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		// fileName = path.Base(file)
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}
