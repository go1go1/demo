package main

import (
	"bufio"
	"demo/tools"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	test := []string{"test"}
	run(test)
}

func run(args []string) {
	if len(args) == 0 {
		tools.Error("please input project name")
		os.Exit(1)
	}

	fileName := args[0]

	in, err := os.Open("./handler.txt")
	if err != nil {
		fmt.Println("open file fail:", err)
		os.Exit(-1)
	}
	defer in.Close()

	// 获取当前目录
	dir, err := os.Getwd()
	tools.CheckIfError(err)
	tools.Info(dir + "/" + fileName + ".go")
	out, err := os.OpenFile(dir+"/"+fileName+".go", os.O_RDWR|os.O_CREATE, 0766)
	tools.CheckIfError(err)

	defer out.Close()

	br := bufio.NewReader(in)
	index := 1
	newName := Capitalize(fileName)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		tools.CheckIfError(err)

		newLine := strings.Replace(string(line), "<name>", newName, -1)
		_, err = out.WriteString(newLine + "\n")
		tools.CheckIfError(err)

		fmt.Println("done ", index)
		index++
	}
	fmt.Println("FINISH!")

}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { //
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
