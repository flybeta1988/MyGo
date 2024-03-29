package main

import (
	"fmt"
	"math"
)

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.3.md

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	//一个接口类型的变量
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	fmt.Printf("The type of areaIntf is: %T\n", sq1)

	areaIntf = sq1

	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}

	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}
