package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["answer"] = 100
	fmt.Println(m)

	m["answer"] = 200
	fmt.Println(m)

	delete(m, "answer")
	fmt.Println(m)

	v, ok := m["answer"]
	fmt.Println("The value:", v, " Present?", ok)
}

