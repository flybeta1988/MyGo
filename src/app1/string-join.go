package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := join("a", "b", "c")
	fmt.Println(str)
}

func join(strs ...string) string {

	//定义一个字节缓冲，快速地连接字符串
	var b bytes.Buffer

	//遍历可变参数列表strs, 类型为[]string
	for _, s := range strs {

		//将遍历的字符串连接写入字节数组
		b.WriteString(s)
	}

	//将连接好的字节数组转换为字符串输出
	return b.String()
}
