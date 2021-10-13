package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c  tac"

	//解析正则表达式，如果成功则返回解释器
	reg1 := regexp.MustCompile(`a\dc`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}

	//result1 := reg1.FindAllStringSubmatch(buf, -1)
	result1 := reg1.FindAllString(buf, -1)
	fmt.Println("result:", result1)

}
