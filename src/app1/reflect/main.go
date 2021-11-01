package main

import (
	"fmt"
	"reflect"
)

type UserModel struct {
	table string
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	password int `json:"password"`
}

func main() {
	//test1()
	//test2()
	//test3()
	test4()
	//test5()
}

func test5() {

	a := 1
	v := reflect.ValueOf(a)

	fmt.Println("v Type:", v.Type())
	fmt.Println("v CanSet:", v.CanSet())

	v = reflect.ValueOf(&a)
	fmt.Println("v Type:", v.Type())
	fmt.Println("v CanSet:", v.CanSet())

	v = v.Elem() // element value
	fmt.Println("v Type:", v.Type())
	fmt.Println("v CanSet:", v.CanSet())

	// set
	v.SetInt(2)
	fmt.Println("after set, v:", v)

	newValue := reflect.ValueOf(3)
	v.Set(newValue)
	fmt.Println("after set, v:", v)
}

func resetStructFields(user *User) {

	v := reflect.ValueOf(user)
	v = v.Elem()

	for i := 0; i < v.NumField(); i ++ {
		//fmt.Println(t.Field(i).Name, v.Field(i).Interface())
		if !v.Field(i).CanSet() {
			continue
		}
		fieldType := v.Field(i).Type().Kind()
		switch fieldType {
		case reflect.Int:
			v.Field(i).SetInt(200)
		case reflect.String:
			v.Field(i).SetString("flyalpha")
		}
		fmt.Println("v type:", fieldType, " CanSet:",  v.Field(i).CanSet())
	}
}

func resetStructFieldsV2(user *User) {

	var fields []interface{}

	v := reflect.ValueOf(user)
	v = v.Elem()

	for i := 0; i < v.NumField(); i ++ {
		//fmt.Println(t.Field(i).Name, v.Field(i).Interface())
		if !v.Field(i).CanSet() {
			continue
		}
		fields = append(fields, v.Field(i))
		/*fieldType := v.Field(i).Type().Kind()
		switch fieldType {
		case reflect.Int:
			v.Field(i).SetInt(200)
		case reflect.String:
			v.Field(i).SetString("flyalpha")
		}
		fmt.Println("v type:", fieldType, " CanSet:",  v.Field(i).CanSet())*/
	}
}

func test4() {
	//user := User{Id:100, Name:"flybeta", Age:33}
	user := User{}
	resetStructFields(&user)
	fmt.Println(user)
}

func test3() {
	typeOfUser := reflect.TypeOf(User{})
	//valueOfUser := reflect.ValueOf(User{})

	if userId, ok := typeOfUser.FieldByName("Id"); ok {
		fmt.Println(userId)
	}
	if userName, ok := typeOfUser.FieldByName("Name"); ok {
		fmt.Println(userName)
	}
	if userAge, ok := typeOfUser.FieldByName("age"); ok {
		fmt.Println(userAge)
	}
	if userMoel, ok := typeOfUser.FieldByName("model"); ok {
		fmt.Println(userMoel)
		fmt.Println(userMoel.Tag.Get("json"))
		fmt.Println(userMoel.Name)
	}
}

func test2() {
	var um = UserModel{"org_course"}
	typeOfA := reflect.TypeOf(um)
	valueOfA := reflect.ValueOf(um)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	fmt.Println(valueOfA.Type(), valueOfA.Kind(), valueOfA.FieldByName("table"))
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
