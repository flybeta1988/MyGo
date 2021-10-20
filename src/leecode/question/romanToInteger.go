package question

import "fmt"

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}

func TestRomanToInteger() {
	charNumMap := map[string]int{
		"III": 3,
		"IV": 4,
		"V": 5,
		"VI": 6,
	}
	for k, v := range charNumMap {
		if v != romanToInt(k) {
			fmt.Println("Error: k:", k, " v:", v)
		} else {
			fmt.Println("Ok: k:", k, " v:", v)
		}
	}
}
