package interview

import "fmt"

func TestConst()  {
	const (
		x = iota
		y
		z = "zz"
		k
		p = iota
	)
	fmt.Println(x, y, z, k, p)
}

func GetValue(m map[int]string, id int) (string, bool) {
	if _, ok := m[id]; ok {
		return "exist data", true
	}
	return "", false
}

func TestReturn() {
	intMap := map[int]string{
		1 : "a",
		2 : "b",
		3 : "c",
	}
	v, err := GetValue(intMap, 3)
	fmt.Println(v, err)
}
