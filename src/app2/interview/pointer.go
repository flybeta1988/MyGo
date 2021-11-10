package interview

import "fmt"

type student struct {
	Name string
	Age  int
}

func PaseStudent() {

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
