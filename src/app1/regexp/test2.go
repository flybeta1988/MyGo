package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 原生字符串
	buf := getHtml()

	//解析正则表达式，如果成功则返回解释器
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	if reg == nil {
		fmt.Println("regexp err")
		return
	}

	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Printf("result type: %T\n", result)

	for _, text := range result {
		fmt.Println("text[1] =", text[1])
	}
}

func getHtml() string {
	return `
	<!DOCTYPE html>
	<html lang="zh-CN">
	<head>
		<title>C语言中文网 | Go语言入门教程</title>
	</head>
	<body>
		<div>Go语言简介</div>
		<div>Go语言基本语法
		Go语言变量的声明
		Go语言教程简明版
		</div>
		<div>Go语言容器</div>
		<div>Go语言函数</div>
	</body>
	</html>
		`
}
