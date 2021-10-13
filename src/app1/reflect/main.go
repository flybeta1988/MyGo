package main

import (
	"fmt"
	"reflect"
)

type UserModel struct {
	table string
}

func main() {
	x := 3.14
	v := reflect.ValueOf(x)
	dType := v.Type()
	v.Kind()

	/*userModel := &UserModel{"user"}
	dType := reflect.TypeOf(userModel)
	dValue := reflect.ValueOf(userModel)*/
	fmt.Println(dType)
	fmt.Println(v)
}
