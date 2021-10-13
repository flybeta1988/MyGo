package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat2 float64

func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex3 struct {
	X, Y float64
}

func (v *Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := Vertex3{3, 4}
	a = f
	a = &v
	fmt.Println(a.Abs())
}