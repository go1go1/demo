package main

import (
	"fmt"
	"strconv"

	hashids "github.com/speps/go-hashids"
)

type test struct {
	Userid string `bson:"userid" json:"userid"`
	Name   string `bson:"name" json:"name"`
	Alias  string `json:"-"`
}

//EncryptID 混淆
func EncryptID(data uint32, salt string) (string, error) {
	if 0 == data {
		return "", nil
	}

	// ID一定是数字
	dataStr := fmt.Sprintf("%0*v", 10, data-10000)

	hd := hashids.NewData()
	hd.Salt = salt
	h, err := hashids.NewWithData(hd)
	if nil != err {
		return "", err
	}

	e, err := h.EncodeHex(dataStr)
	if nil != err {
		return "", err
	}
	return e, nil
}

//DecryptID 还原混淆
func DecryptID(data string, salt string) (uint32, error) {
	if "" == data {
		return 0, nil
	}

	hd := hashids.NewData()
	hd.Salt = salt
	h, err := hashids.NewWithData(hd)
	if nil != err {
		return 0, err
	}

	e, err := h.DecodeHex(data)
	if nil != err {
		return 0, err
	}

	// 必须还原到Uint格式
	dataNum, err := strconv.ParseUint(e, 10, 32)
	return uint32(dataNum + 10000), nil
}

func main() {
	//str := "{\"userid\": \"zhangsan\",\"name\": \"张三\",\"alias\": \"jackzhang\"}"

	//t := &test{}
	//err := json.Unmarshal([]byte(str), t)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//fmt.Printf("%#v\n", t)
	//to, _ := json.Marshal(t)

	//strSli := []int32("RYFK3")
	//var res uint32 = 0
	//fmt.Printf("%#v\n", string(rune(51)))
	//for i := 0; i < len(strSli); i++ {
	//	res += uint32(strSli[i])
	//}
	//id, _ := EncryptID(22341, "corp")
	//fmt.Printf("%v\n", id)

	num, _ := DecryptID("zETyT7T3TVTkTETMTlSg", "corp")
	fmt.Printf("%v\n", num)

}
