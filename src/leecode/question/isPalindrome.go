package question

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if (x < 0) {
		return false
	}
	str := strconv.Itoa(x)
	size := len(str)
	for i := 1; i <= size/2; i++ {
		if (str[i-1] != str[size-i]) {
			return false
		}
	}
	return true
}

func isPalindromeV2(x int) bool {
	if (x < 0) {
		return false
	}
	t := x
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x /= 10
		fmt.Printf("x:%d, y:%d\n", x, y)
	}
	return t == y
}

func isPalindromeV3(x int) bool {
	if (x < 0 || (x != 0 && x%10 == 0)) {
		return false
	}
	y := 0
	for x > y {
		y = y*10 + x%10
		x /= 10
		fmt.Printf("x:%d, y:%d\n", x, y)
	}
	return x == y || x == y/10
}

func TestIsPalindrome() {
	x := 123454321
	//fmt.Printf("i:%d, j:%d, s1:%s, s2:%s\n", i-1, j, string(str[i-1]), string(str[j]))
	if (isPalindromeV3(x)) {
		fmt.Println(x, "是回文数")
	} else {
		fmt.Println(x, "不是回文数")
	}
}
