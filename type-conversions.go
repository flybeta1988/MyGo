package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y = 3, 4
	var e float64 = float64(x*y + y*y)
	var f float64 = math.Sqrt(e)
	var z int = int(f)
	fmt.Println(x, y, z, e, f)
}
