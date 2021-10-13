package main

import (
	"fmt"
)

func main() {
	//defer_call()
	//pase_student()
	testExtend()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func testExtend() {
	t := Teacher{}
	t.ShowA()
	t.ShowB()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {

	students := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	fmt.Println(students)

	m := make(map[string]*student)
	for i := 0; i < len(students); i ++ {
		m[students[i].Name] = &students[i]
	}

	for k, v := range m {
		fmt.Println(k, "=", v)
	}
	//fmt.Println(m)
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}