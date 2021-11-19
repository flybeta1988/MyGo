package main

import (
	"encoding/json"
	"fmt"
)

type Card struct {
	ID int
	Type int
	Name string
}

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int
	Cards []Card
}

type JsonTest struct {}

func main() {
	jt := &JsonTest{}
	jt.test1()
	//jt.test2()
}

func (j JsonTest) test1() {
	user := []User{{ID: 100, Name: "abc", Age: 30}, {ID: 101, Name: "efg", Age: 31}}
	//user := User{ID: 100, Name: "abc", Age: 30}
	//fmt.Println(user)
	result, err := json.Marshal(user)
	if err != nil {

	}
	r := string(result)
	fmt.Println(r)
}

func (j JsonTest) test2() {
	str := "{\"ID\":100,\"Name\":\"abc\",\"Age\":30,\"Cards\":null}"
	s := []byte(str)

	var result User
	err := json.Unmarshal(s, &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
