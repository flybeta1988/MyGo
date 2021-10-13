package main

import "fmt"

type DataWriter interface {
	WriterLog(data interface{}) error
	WriterData(data interface{}) error
}

type file struct {}

// 实现DataWriter接口的WriteData方法
func (d *file) WriterLog(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

// 实现DataWriter接口的WriteData方法
func (d *file) WriterData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func main() {
	// 实例化file
	f := new(file)

	// 声明一个DataWriter接口
	var writer DataWriter

	// 将接口赋值f，也就是*file类型
	writer = f

	writer.WriterData("this is test data")
}
