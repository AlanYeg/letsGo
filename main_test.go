package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestHelloWorld(t *testing.T) {
	t.Log("hello wo       rld")
	t.Log("hello wo       rld")
	t.Log("hello wo       rld")
	t.Log("hello wo       rld")

}

func TestT(t *testing.T) {
	i := "123"
	fmt.Printf("Size of var (reflect.TypeOf.Size): %d\n", reflect.TypeOf(i).Size())
	fmt.Printf("Size of var (unsafe.Sizeof): %d\n", unsafe.Sizeof(i[0]))

}
func TestX(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")
	file, err := os.Create("D:/kaka.txt")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(file, r); err != nil {
		log.Fatal(err)
	}
}

func TestX1(t *testing.T) {
	fmt.Printf("|%06d|%-06x|\n", 12, 345)

}

func TestBinaryFile_PrintTo(t *testing.T) {
	var test BinaryFile
	test.data = "asdasd各种牛人，膜拜 The acme of foolishness这边的文章都不错"
	test.printType = "x"
	test.PrintTo(os.Stdout)
}
