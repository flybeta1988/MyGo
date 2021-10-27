package main

import (
	"fmt"
	"reflect"
)

type UserModel struct {
	table string
}

func main() {
	test2()
}

func test2() {
	var um UserModel
	typeOfA := reflect.TypeOf(um)
	valueOfA := reflect.ValueOf(um)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	fmt.Println(valueOfA.Type(), valueOfA.Kind())
}

func test1() {
	x := 3.14
	v := reflect.ValueOf(x)
	dType := v.Type()
	kind := v.Kind()
	fmt.Println(kind)

	/*userModel := &UserModel{"user"}
	dType := reflect.TypeOf(userModel)
	dValue := reflect.ValueOf(userModel)*/
	fmt.Println(dType)
	fmt.Println(v)
}
