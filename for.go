package main

import "fmt"

func for1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func for2() {
	sum := 1
	for ; sum < 100 ; {
		sum += sum
	}
	fmt.Println(sum)
}

func for3() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	for1()
	for2()
	for3()
}
