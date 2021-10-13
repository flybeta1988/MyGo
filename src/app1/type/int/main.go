package main

import (
	"fmt"
	"math"
)

func main() {
	testBase()
}

func testBase() {
	fmt.Printf("int8 value range:%d~%d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16 value range:%d~%d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32 value range:%d~%d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64 value range:%d~%d\n", math.MinInt64, math.MaxInt64)
}
