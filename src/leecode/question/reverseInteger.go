package question

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	var y int
	var result string
	str := strconv.Itoa(x)
	size := len(str)
	for i := size - 1; i >= 0; i-- {
		result += (string(str[i]))
	}
	//fmt.Println(result)
	y, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}
	return y
}

func reverseV2(x int) int {
	y := 0
	for x != 0 {
		y = y * 10 + x % 10
		if !(-(1 << 31) <= y && y <= (1 << 31) - 1) {
			return 0
		}
		x /= 10
	}
	return y
}

func reverseV3(x int) int {
	y := -(1 << 2)
	fmt.Println(y)
	return y
}

func TestReverse() {
	result := reverseV2(123456)
	fmt.Println(result)
}
