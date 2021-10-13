package main

import "fmt"

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e2
	fmt.Println(v.X)
}
