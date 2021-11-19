package interview

import "fmt"

func DeferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func DeferCallFun() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

//答案：https://studygolang.com/articles/10746
//https://my.oschina.net/henrylee2cn/blog/505535
func TestDeferFunc() {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		fmt.Println("defer t:", t)
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		fmt.Println("defer t:", t)
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		fmt.Println("defer t:", t)
		t += i
	}()
	return 2
}