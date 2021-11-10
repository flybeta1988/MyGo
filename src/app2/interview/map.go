package interview

import "fmt"

//考点：make默认值和append
func TestAppend() {
	s := make([]int, 5)
	//s := make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //[0 0 0 0 0 1 2 3]
}
