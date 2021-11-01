package main

import (
	"fmt"
	"strconv"
)

func main() {
	test1()
}

func test1() {
	str := "中国"
	strUnicode := strconv.QuoteToASCII(str)
	fmt.Println(strUnicode)
}
