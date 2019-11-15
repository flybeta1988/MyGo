package main

import "fmt"

func main() {

	var a []int
	printSlices("a", a)

	a = append(a, 0)
	printSlices("a", a)

	a = append(a, 1)
	printSlices("a", a)

	a = append(a, 2, 3, 4)
	printSlices("a", a)

}

func printSlices(s string, x []int)  {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
