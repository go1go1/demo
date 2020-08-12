/**
 * Author: richen
 * Date: 2020-08-11 09:48:33
 * LastEditTime: 2020-08-11 20:30:10
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package logi

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	fileNameFormatStr    = "%s/%s_%s.log"
	fileNameFormatErrStr = "%s/%s_%s.err.log"
)

type FileLog struct {
	level         int
	logPath       string
	logName       string
	file          *os.File
	errFile       *os.File
	logDataChan   chan *LogData
	logSplitType  int
	logSplitSize  int64
	lastSplitHour int
}

func NewFileLog(config map[string]string) (err error, log LogiInterface) {
	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("Not found log_name")
		return
	}
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("Not found log_path")
		return
	}
	level, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("Not found log_level")
		return
	}

	var (
		logSplitType int
		logSplitSize int64
	)
	splitStr, ok := config["log_split_type"]
	if !ok {
		splitStr = "hour"
	}
	if splitStr == "size" {
		splitSizeStr, ok := config["log_split_size"]
		if !ok {
			splitSizeStr = "104857600"
		}
		logSplitSize, err = strconv.ParseInt(splitSizeStr, 10, 64)
		if err != nil {
			logSplitSize = 104857600
		}
		logSplitType = LogSplitTypeSize
	} else {
		logSplitType = LogSplitTypeHour
	}

	chanSize, ok := config["log_chan_size"]
	if !ok {
		chanSize = "50000"
	}

	cSize, err := strconv.Atoi(chanSize)
	if err != nil {
		cSize = 50000
	}

	log = &FileLog{
		// level:   level,
		logName:       logName,
		logPath:       logPath,
		logDataChan:   make(chan *LogData, cSize),
		logSplitType:  logSplitType,
		logSplitSize:  logSplitSize,
		lastSplitHour: time.Now().Hour(),
	}
	log.SetLevel(getLevel(level))
	log.init()
	return
}

func (f *FileLog) init() {
	date := time.Now().Format("2006-01-02")

	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Open log file dir %s failed, err: %#v", dir, err))
	}
	f.logPath = dir + "/" + f.logPath
	if !exists(f.logPath) {
		mkDir(f.logPath)
	}

	filename := fmt.Sprintf(fileNameFormatStr, f.logPath, f.logName, date)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("Open log file %s failed, err: %#v", filename, err))
	}
	f.file = file

	filename = fmt.Sprintf(fileNameFormatErrStr, f.logPath, f.logName, date)
	errFile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("Open errlog file %s failed, err: %#v", filename, err))
	}
	f.errFile = errFile

	go f.writeLog()
}

func (f *FileLog) checkSplitFile(errFile bool) {
	if f.logSplitType == LogSplitTypeHour {
		f.splitHour(errFile)
	} else {
		f.splitSize(errFile)
	}
}

func (f *FileLog) splitHour(errFile bool) {
	now := time.Now()
	if now.Hour() == f.lastSplitHour {
		return
	}
	var filename, backupFileName string
	if errFile {
		backupFileName = fmt.Sprintf(fileNameFormatErrStr, f.logPath, f.logName, now.Format("2006-01-02")+"-"+fmt.Sprintf("%02d", f.lastSplitHour))
	} else {
		backupFileName = fmt.Sprintf(fileNameFormatStr, f.logPath, f.logName, now.Format("2006-01-02")+"-"+fmt.Sprintf("%02d", f.lastSplitHour))
	}

	file := f.file
	if errFile {
		file = f.errFile
		filename = fmt.Sprintf(fileNameFormatErrStr, f.logPath, f.logName, now.Format("2006-01-02-15"))
	} else {
		filename = fmt.Sprintf(fileNameFormatStr, f.logPath, f.logName, now.Format("2006-01-02-15"))
	}

	file.Close()
	err := os.Rename(filename, backupFileName)
	if err != nil {
		_ = fmt.Errorf("Log file backup faild.")
		return
	}
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		_ = fmt.Errorf("Log file reopen faild.")
		return
	}
	if errFile {
		f.errFile = file
	} else {
		f.file = file
	}
	f.lastSplitHour = now.Hour()
}

func (f *FileLog) splitSize(errFile bool) {
	file := f.file
	if errFile {
		file = f.errFile
	}

	fileSize := fSize(file)
	fmt.Printf("%d, %d, %s \n", fileSize, f.logSplitSize, f.logPath)
	if fileSize < f.logSplitSize {
		return
	}
	now := time.Now()
	var filename, backupFileName string
	if errFile {
		backupFileName = fmt.Sprintf(fileNameFormatErrStr, f.logPath, f.logName, now.Format("2006-01-02-15-04-05"))
	} else {
		backupFileName = fmt.Sprintf(fileNameFormatStr, f.logPath, f.logName, now.Format("2006-01-02-15-04-05"))
	}

	if errFile {
		filename = fmt.Sprintf(fileNameFormatErrStr, f.logPath, f.logName, now.Format("2006-01-02"))
	} else {
		filename = fmt.Sprintf(fileNameFormatStr, f.logPath, f.logName, now.Format("2006-01-02"))
	}

	fmt.Printf("%d, %d, %s, %s \n", fileSize, f.logSplitSize, backupFileName, filename)

	err := file.Close()
	err = os.Rename(filename, backupFileName)
	if err != nil {
		_ = fmt.Errorf("Log file backup faild.")
		return
	}
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("Log file reopen faild: %#v", err)
		return
	}
	if errFile {
		f.errFile = file
	} else {
		f.file = file
	}

}

// printLog
func (f *FileLog) printLog(color bool, level int, format string, args ...interface{}) {
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
	logData := &LogData{
		Message:  msg,
		TimeStr:  now.Format("2006-01-02 15:04:05.999"),
		ColorStr: colorStr,
		LevelStr: getLevelText(level),
		ResetStr: resetStr,
	}
	if level >= LogLevelWarn {
		fileName, funcName, lineNo = getLineInfo()
		logData.FileName = fileName
		logData.FuncName = funcName
		logData.LineNo = lineNo
	}
	select {
	case f.logDataChan <- logData:
	default:
	}
}

// writeLog
func (f *FileLog) writeLog() {
	for data := range f.logDataChan {
		if data.LineNo > 0 {
			f.checkSplitFile(true)
			fmt.Fprintf(f.errFile, "%s[%s] [%s] %s \n%s %s %d\n%s", data.ColorStr, data.TimeStr, data.LevelStr, data.Message, data.FileName, data.FuncName, data.LineNo, data.ResetStr)
		} else {
			f.checkSplitFile(false)
			fmt.Fprintf(f.file, "%s[%s] [%s] %s\n%s", data.ColorStr, data.TimeStr, data.LevelStr, data.Message, data.ResetStr)
		}
	}
}

func (f *FileLog) Close() {
	f.file.Close()
	f.errFile.Close()
}

func (f *FileLog) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *FileLog) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	f.printLog(false, LogLevelDebug, format, args...)
}
func (f *FileLog) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	f.printLog(false, LogLevelTrace, format, args...)
}
func (f *FileLog) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	f.printLog(false, LogLevelInfo, format, args...)
}
func (f *FileLog) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	f.printLog(false, LogLevelWarn, format, args...)
}
func (f *FileLog) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	f.printLog(false, LogLevelError, format, args...)
}
func (f *FileLog) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	f.printLog(false, LogLevelFatal, format, args...)
}
