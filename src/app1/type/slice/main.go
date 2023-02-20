package main

import "fmt"

//数组切片的`数据结构`可以抽象为以下3个变量：
//一个指向原生数组的指针
//数组切片中的元素个数
//数组切片已分配的存储空间
func main() {
	//testBase()
	//testMake()
	//testOpt()
	testCap()
}

func testCap() {
	vals := make([]int, 5)
	for i := 1; i <= 3; i ++ {
		vals = append(vals, i)
	}
	fmt.Println(len(vals), cap(vals))
	fmt.Println(vals)
}

func testMake() {
	//idList := make([]int, 5)
	//idList := make([]int, 5, 10)
	idList := []int{1, 2, 3, 4, 5}
	fmt.Println(len(idList))
}

func testOpt() {
	var idList []int
	printSlices("idList", idList)

	idList = append(idList, 1)
	printSlices("idList", idList)

	idList = append(idList, 2, 3)
	printSlices("idList", idList)
}

func printSlices(varName string, x []int)  {
	fmt.Printf("%s len=%d cap=%d %v\n", varName, len(x), cap(x), x)
}

func testBase() {
	var idList [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sliceIdList1 []int = idList[:]
	var sliceIdList2 []int = idList[:5]
	var sliceIdList3 []int = idList[1:]
	fmt.Println(sliceIdList1)
	fmt.Println(sliceIdList2)
	fmt.Println(sliceIdList3)
}
