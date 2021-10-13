package main

import "fmt"

type Rect struct {
	x, y float64
	width, height float64
}

func main() {
	testBase()
}

func NewRect(x, y, width, height float64) *Rect {
	/*rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}*/
	return &Rect{x, y, width, height}
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func testBase() {
	var rect *Rect
	rect = NewRect(0,0, 100, 200)
	fmt.Println(rect)
}