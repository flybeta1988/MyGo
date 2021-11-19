package interview

import "fmt"

//考点：make默认值和append
func TestAppend() {
	s := make([]int, 5)
	//s := make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //[0 0 0 0 0 1 2 3]
}

func TestAppend2() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func TestNew() {
	//list := new([]int)
	//list := make([]int, 0)
	var list []int
	list = append(list, 1)
	fmt.Println(list)
}
