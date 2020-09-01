package protodata

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

/**
protobuf
IDL编写
生成指定语言代码
编码解码

基础类型IDL都有对应的定义

其他类型：
枚举:
enum EnumAllow {
	STARTED = 0
	ENDED   = 1
}

结构体:
message Person {
	//后面的数字表示标识号，并不是值
	int32 id = 1;
	string name = 2;
	//repeated表示可重复(数组，go里面是切片)
	repeated Phone phones = 3; //可以有多个手机
}


*/

func writeProtoContent(filename string) (err error) {
	var contactBook ContactBook
	var i int32
	for i = 0; i < 100; i++ {
		p := &Person{
			Id:   i,
			Name: fmt.Sprintf("a%d", i),
		}
		phone := &Phone{
			Type:   PhoneType_HOME,
			Number: "13333333333",
		}

		p.Phones = append(p.Phones, phone)
		contactBook.Persons = append(contactBook.Persons, p)
	}

	data, err := proto.Marshal(&contactBook)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
	return
}

func readProtoContent(filename string) (err error) {
	var contactBook ContactBook
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = proto.Unmarshal(data, &contactBook)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	return
}
