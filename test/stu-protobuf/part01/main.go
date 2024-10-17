// Package main
// @Author twilikiss 2024/4/26 19:55:55
package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	"part01/out"
)

func main() {
	test := &out.Student{
		Name:   "elysia",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &out.Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}
}
