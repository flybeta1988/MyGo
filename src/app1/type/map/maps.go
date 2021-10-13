package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m1 map[string]Vertex

func main() {
	m1 = make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

