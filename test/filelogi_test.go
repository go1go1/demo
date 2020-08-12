/**
 * Author: richen
 * Date: 2020-08-11 10:24:25
 * LastEditTime: 2020-08-11 16:09:33
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package test

import (
	"demo/logi"
	"testing"
	"time"
)

var config map[string]string

func TestFileLogger(t *testing.T) {
	config = make(map[string]string)
	config["log_path"] = "./logs"
	config["log_name"] = "test"
	config["log_level"] = "debug"

	_, logger := logi.NewFileLog(config)
	defer logger.Close()
	logger.Debug("user id[%d]", 2332423)
	logger.Trace("user name %s", "sam")
	logger.Info("user sex %s", "男")
	logger.Warn("user age %d", 18)
	logger.Error("user error %s", "error")
	logger.Fatal("user fatal %s", "fatal")
	time.Sleep(5 * time.Second)
}

func TestConsoleLogger(t *testing.T) {
	config = make(map[string]string)
	config["log_level"] = "debug"
	_, logger := logi.NewConsoleLog(config)
	defer logger.Close()
	logger.Debug("user id[%d]", 2332423)
	logger.Trace("user name %s", "sam")
	logger.Info("user sex %s", "男")
	logger.Warn("user age %d", 18)
	logger.Error("user error %s", "error")
	logger.Fatal("user fatal %s", "fatal")
	time.Sleep(2 * time.Second)
}
