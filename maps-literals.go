package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	var m = map[string]Vertex {
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			44.68433, -84.39967,
		},
	}
	fmt.Println(m)
}

