package util

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"sync"
	"testing"
)

type Test struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	sync.Mutex
}

func TestEncode(t *testing.T) {
	test := &Test{
		Id:   1,
		Name: "tommmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm",
	}
	b, err := Encode(test)
	if nil != err {
		t.Errorf("TestEncode err : %+v", err)
	}
	fmt.Println(len(b))

	var test2 Test
	err = Decode(b, &test2)
	if nil != err {
		t.Errorf("TestEncode err : %+v", err)
	}
	fmt.Println(test2)
}

func TestEncodeGZip(t *testing.T) {
	test := &Test{
		Id:   1,
		Name: "tommmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm",
	}
	b, err := EncodeGZip(test)

	if nil != err {
		t.Errorf("TestEncodeGZip error, %+v", err)
	}
	fmt.Println(len(b))
	var test2 Test
	err = DecodeUnGzip(b, &test2)
	if nil != err {
		t.Errorf("TestEncodeGZip error, %+v", err)
	}
	fmt.Println(test2)
}

func TestGZip(t *testing.T) {
	b := []byte("hello")
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, err := zw.Write(b)
	if err != nil {
		t.Error(err)
	}

	if err := zw.Close(); err != nil {
		t.Error(err)
	}

	b2 := buf.Bytes()
	fmt.Printf("b2 : %+v \n", b2)

}
