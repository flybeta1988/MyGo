package main

import "fmt"

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 0}
	v3 = Vertex{}
	p = &Vertex{1, 2}
)
func main() {
	fmt.Println(v1, p, v2, v3)
}
