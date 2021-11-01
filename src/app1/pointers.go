package main

import "fmt"

func main() {

	i, j := 42, 2701

	p := &i //&取址
	fmt.Println(*p)

	*p = 21 //*取值
	fmt.Println(i)

	p = &j
	*p = *p
	fmt.Println(j)

}
