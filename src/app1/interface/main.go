package main

import (
	pkgOne "app1/package1"
	pkgTwo "app1/package2"
	"fmt"
)

type File struct {}

type IFile interface {
	Read(buf []byte) (n int)
	//Write(buf []byte) (n int)
}
type IReader interface {
	Read(buf []byte) (n int)
}
type IWriter interface {
	Write(buf []byte) (n int)
}

func main() {
	//var file1 IFile = new(File)
	//testAssign()
	testQuery()
}

func (f *File) Read(buf []byte) (n int) {
	return 0
}

func (f *File) Write(buf []byte) (n int) {
	return 0
}

func testAny() {
	/*var v1 interface{} = 1 //将int类型赋值给interface{}
	var v2 interface{} = "abc" //将string类型赋值给interface{}
	var v3 interface{} = &v2 //将*interface{}类型赋值给interface{}
	var v4 interface{} = struct {X int}{1}
	var v5 interface{} = &v4*/
}

func testQuery()  {
	var file1 IWriter = new(File)

	//IWriter 接口转换为 two.IStream 接口
	if file5, ok := file1.(pkgTwo.IStream); ok {
		fmt.Println("file5 convert ok", file5)
	} else {
		fmt.Println("file5 convert fail!")
	}

	//判断接口指向的对象是否是某个类型
	if file6, ok := file1.(*File); ok {
		fmt.Println("file6 convert ok", file6)
	} else {
		fmt.Println("file6 convert ok")
	}
}

func testAssign()  {
	var file1 pkgTwo.IStream = new(File)
	var file2 pkgOne.ReadWriter = file1
	var file3 pkgTwo.IStream = file2
	var file4 IWriter = file1
	//var file5 pkgTwo.IStream = file4
	fmt.Println(file3, file4)
}
