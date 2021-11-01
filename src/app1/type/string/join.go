package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

//来源：https://blog.csdn.net/hatlonely/article/details/79156311
func main() {

}

//直接使用运算符
func BenchmarkAddStringWithOperator(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = hello + "," + world
	}
}

func BenchmarkAddMoreStringWithOperator(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		var str string
		for i := 0; i < 100; i++ {
			str += hello + "," + world
		}
	}
}

//fmt.Sprintf()
func BenchmarkAddStringWithSprintf(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s,%s", hello, world)
	}
}

//strings.Join
func BenchmarkAddStringWithJoin(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{hello, world}, ",")
	}
}

//buffer.WriteString()
func BenchmarkAddStringWithBuffer(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(hello)
		buffer.WriteString(",")
		buffer.WriteString(world)
		_ = buffer.String()
	}
}

func BenchmarkAddMoreStringWithBuffer(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		for i := 0; i < 100; i++ {
			buffer.WriteString(hello)
			buffer.WriteString(",")
			buffer.WriteString(world)
		}
		_ = buffer.String()
	}
}



