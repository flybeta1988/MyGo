package question

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {

	size := len(strs)
	if 1 == size {
		return strs[0]
	}

	var nextIndex int
	var prefix, nextStr string

	for i, str := range strs {
		nextIndex = i + 1
		if nextIndex == size {
			break
		}

		if i > 0 {
			str = prefix
			prefix = ""
		}

		nextStr = strs[nextIndex]
		//fmt.Printf("i:%d str:%s nextStr:%s\n", i, str, nextStr)

		for j, s := range str {
			if j > len(nextStr)-1 || string(s) != string(nextStr[j]) {
				break
			}
			prefix = prefix + string(s)
		}
	}
	return prefix
}

func TestLongestCommonPrefix() {
	//var strs = []string{"flower", "flow", "flight"}
	//var strs = []string{"a"}
	//var strs = []string{"cir","car"}
	//var strs = []string{"ab","a"}

	var strs_list [][]string
	strs_list = append(strs_list, []string{"flower", "flow", "flight"})
	strs_list = append(strs_list, []string{"a"})
	strs_list = append(strs_list, []string{"cir","car"})
	strs_list = append(strs_list, []string{"ab","a"})
	//fmt.Println(strs)

	for _, strs := range strs_list {
		result := longestCommonPrefix(strs)
		fmt.Println(result)
	}
}
