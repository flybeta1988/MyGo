package question

import "fmt"

func TestTwoSum() {
	var nums []int
	nums = []int{11, 15, 2, 7}
	fmt.Println(twoSum3(nums, 9))
}

func twoSum(nums []int, target int) []int {
	var indexes []int
	if 0 == len(nums) {
		return indexes
	}

	size := len(nums)
	result := make([]int, 0)
	for i := 0; i < size; i++ {
		for j := 1; j < size; j++ {
			if (nums[i] + nums[j] == target) {
				//fmt.Printf("i:%d + j:%d = %d\n", nums[i], nums[j], target)
				return []int{i, j}
			}
		}
	}
	return result
}

func twoSum2(nums []int, target int) []int {
	for k1, _ := range nums {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			if target == nums[k1] + nums[k2] {
				return []int{k1, k2}
			}
		}
	}
	return []int{}
}

//使用map方式解题，降低时间复杂度
func twoSum3(nums []int, target int) []int {
	m := make(map[int]int)
	for index, value := range nums {
		if preIndex, ok := m[target -value]; ok {
			return []int{preIndex, index}
		} else {
			m[value] = index
		}
	}
	return []int{}
}