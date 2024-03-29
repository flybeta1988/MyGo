package main

import "fmt"

type Vertex3 struct {
	Lat, Long float64
}

var m map[string]Vertex3

func main() {
	var m = map[string]Vertex3 {
		"Bell Labs": Vertex3{
			40.68433, -74.39967,
		},
		"Google": Vertex3{
			44.68433, -84.39967,
		},
	}
	fmt.Println(m)
}

