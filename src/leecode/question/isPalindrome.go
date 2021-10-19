package question

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	size := len(str)
	for i := 1; i <= size/2; i++ {
		if (str[i-1] != str[size-i]) {
			return false
		}
	}
	return true
}

func TestIsPalindrome() {
	x := 123454321
	//fmt.Printf("i:%d, j:%d, s1:%s, s2:%s\n", i-1, j, string(str[i-1]), string(str[j]))
	if (isPalindrome(x)) {
		fmt.Println(x, "是回文数")
	} else {
		fmt.Println(x, "不是回文数")
	}
}
