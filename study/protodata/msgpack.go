package protodata

import (
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/vmihailenco/msgpack"
)

type Person2 struct {
	Name string `bson:"name" json:"name"`
	Age  int    `bson:"age" json:"age"`
	Sex  string `bson:"sex" json:"sex"`
}

func writeContent(filename string) (err error) {
	var persons []*Person2
	for i := 0; i < 10; i++ {
		p := &Person2{
			Name: fmt.Sprintf("name%d", i),
			Age:  rand.Intn(100),
			Sex:  "Man",
		}

		persons = append(persons, p)
	}

	data, err := msgpack.Marshal(persons)
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

func readContent(filename string) (err error) {
	var parsons []*Person2
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = msgpack.Unmarshal(data, &parsons)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	return
}
