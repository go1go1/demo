package protodata

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	filename := "./person.txt"
	//err := writeContent(filename)
	//if err != nil {
	//	fmt.Printf("%#v\n", err)
	//	return
	//}

	err := readContent(filename)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	t.Logf("success")
}
