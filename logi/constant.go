/**
 * Author: richen
 * Date: 2020-08-11 10:07:13
 * LastEditTime: 2020-08-11 17:57:30
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package logi

const (
	LogLevelDebug = iota
	LogLevelTrace
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

// Split type
const (
	LogSplitTypeHour = iota
	LogSplitTypeSize
)

// getColor
func getColor(level int) string {
	switch level {
	case LogLevelDebug:
		return White
	case LogLevelTrace:
		return Magenta
	case LogLevelInfo:
		return Blue
	case LogLevelWarn:
		return Yellow
	case LogLevelError:
		return Red
	case LogLevelFatal:
		return RedBold
	default:
		return White
	}
}

// getLevelText
func getLevelText(level int) string {
	switch level {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelTrace:
		return "TRACE"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	}
	return "UNKNOWN"
}

// getLevel
func getLevel(level string) int {
	switch level {
	case "debug":
		return LogLevelDebug
	case "trace":
		return LogLevelTrace
	case "info":
		return LogLevelInfo
	case "warn":
		return LogLevelWarn
	case "error":
		return LogLevelError
	case "fatal":
		return LogLevelFatal
	}
	return LogLevelDebug
}
