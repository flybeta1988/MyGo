package main

import (
	"fmt"
	"strconv"
)

//https://blog.csdn.net/shenyunfei1990/article/details/84922931

type Element interface{}

type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{name: "flybeta", age: 32}

	for i, e := range list {
		//方法一
		if v, ok := e.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", i, v)
		} else if v, ok := e.(string); ok {
			fmt.Printf("list[%d] is an string and its value is %s\n", i, v)
		} else if v, ok := e.(Person); ok {
			fmt.Printf("list[%d] is an Person and its value is %s\n", i, v)
		} else {
			fmt.Println("list[%d] is of a different type", i)
		}

		//方法二
		switch v := e.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", i, v)
		case string:
			fmt.Printf("list[%d] is an string and its value is %s\n", i, v)
		case Person:
			fmt.Printf("list[%d] is an Person and its value is %s\n", i, v)
		default:
			fmt.Println("list[%d] is of a different type", i)
		}
	}
}
