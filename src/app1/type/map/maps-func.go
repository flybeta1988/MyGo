package main

import (
	"flag"
	"fmt"
)

func main() {

	var skill = flag.String("skill", "", "skill to perform")

	flag.Parse()

	var skills = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
	}

	f, ok := skills[*skill]
	fmt.Println(ok)
	if ok {
		f()
	} else {
		fmt.Println("skill not found!")
	}
}
