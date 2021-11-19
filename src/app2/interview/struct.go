package interview

import "fmt"

func TestStructCompare() {
	sn1 := struct {
		age int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	/* 结构体属性中有不可以比较的类型，如map,slice
	sm1 := struct {
		age int
		m map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}*/
}