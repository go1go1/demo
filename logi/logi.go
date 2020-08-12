/**
 * Author: richen
 * Date: 2020-08-11 09:54:39
 * LastEditTime: 2020-08-11 16:45:44
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package logi

import "fmt"

type LogiInterface interface {
	init()
	Close()
	SetLevel(level int)
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

var log LogiInterface

// InitLogger
func InitLogger(name string, config map[string]string) (err error) {
	switch name {
	case "file":
		err, log = NewFileLog(config)
	case "console":
		err, log = NewConsoleLog(config)
	default:
		err = fmt.Errorf("Unsupport logger name:%s", name)
	}
	return
}

func Debug(format string, args ...interface{}) {
	log.Debug(format, args...)
}
func Trace(format string, args ...interface{}) {
	log.Trace(format, args...)
}
func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}
func Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}
func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}
func Fatal(format string, args ...interface{}) {
	log.Fatal(format, args...)
}
