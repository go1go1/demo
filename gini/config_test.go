package gini

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type Config struct {
	Server ServerConfig `ini:"server"`
	Mysql  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MysqlConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database string `ini:"database"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func TestIniConfig(t *testing.T) {
	fmt.Println("hello")

	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("read file faild")
	}

	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Errorf("UnMarshal failed. err:%v", err)
		return
	}
	t.Logf("UnMarshal success, conf:%#v", conf)

	confData, err := Marshal(conf)
	if err != nil {
		t.Errorf("Marshal failed. err:%v", err)
		return
	}
	t.Logf("Marshal success, conf:\n%v", string(confData))
}
