/**
 * Author: richen
 * Date: 2020-08-11 10:06:09
 * LastEditTime: 2020-08-11 20:27:27
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package logi

import (
	"fmt"
	"time"
)

type ConsoleLog struct {
	level int
}

func NewConsoleLog(config map[string]string) (err error, log LogiInterface) {
	level, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("Not found log_level")
		return
	}

	log = &ConsoleLog{}
	log.SetLevel(getLevel(level))
	log.init()
	return
}

func (f *ConsoleLog) init() {

}

func (f *ConsoleLog) Close() {}

func (f *ConsoleLog) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

// printLog
func (f *ConsoleLog) printLog(color bool, level int, format string, args ...interface{}) {
	var (
		fileName, funcName string
		lineNo             int
		colorStr           string
		resetStr           string
	)
	now := time.Now()
	if color {
		resetStr = Reset
		colorStr = getColor(level)
	}
	msg := fmt.Sprintf(format, args...)

	if level >= LogLevelWarn {
		fileName, funcName, lineNo = getLineInfo()
		fmt.Printf("%s[%s] [%s] %s \n%s %s %d\n%s", colorStr, now.Format("2006-01-02 15:04:05.999"), getLevelText(level), msg, fileName, funcName, lineNo, resetStr)
	} else {
		fmt.Printf("%s[%s] [%s] %s\n%s", colorStr, now.Format("2006-01-02 15:04:05.999"), getLevelText(level), msg, resetStr)
	}
}

func (f *ConsoleLog) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	f.printLog(true, LogLevelDebug, format, args...)
}
func (f *ConsoleLog) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	f.printLog(true, LogLevelTrace, format, args...)
}
func (f *ConsoleLog) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	f.printLog(true, LogLevelInfo, format, args...)
}
func (f *ConsoleLog) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	f.printLog(true, LogLevelWarn, format, args...)
}
func (f *ConsoleLog) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	f.printLog(true, LogLevelError, format, args...)
}
func (f *ConsoleLog) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	f.printLog(true, LogLevelFatal, format, args...)
}
