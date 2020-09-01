/**
 * Author: richen
 * Date: 2020-08-11 15:16:31
 * LastEditTime: 2020-08-12 18:20:59
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import (
	"fmt"
	"time"

	"Script/logi"
)

func initLog() {
	config := make(map[string]string)
	config["log_path"] = "logs"
	config["log_name"] = "test"
	config["log_level"] = "debug"
	config["log_split_type"] = "size"  // size, hour
	config["log_split_size"] = "10240" // size, hour
	err := logi.InitLogger("file", config)
	if err != nil {
		_ = fmt.Errorf("logger error")
		return
	}

	logi.Debug("Init logger success")
}

func run() {
	for {
		logi.Info("User server is running")
		logi.Trace("trace %s", "sam")
		logi.Info("info %s", "男")
		logi.Warn("warn %d", 18)
		logi.Error("error %s", "error")
		logi.Fatal("fatal %s", "fatal")
		time.Sleep(time.Second)
	}
}

func main() {
	initLog()
	run()
}
