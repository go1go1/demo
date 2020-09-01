package protodata

import (
	"fmt"
	"testing"
)

func TestRunProto(t *testing.T) {
	filename := "./person_proto.txt"
	err := writeProtoContent(filename)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}

	err = readProtoContent(filename)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	t.Logf("success")
}
