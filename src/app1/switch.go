package main

import (
	"fmt"
	"runtime"
	"time"
)

func switch1()  {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("%s", os)
	}
}

//没有条件的 switch 同 `switch true` 一样
//这一构造使得可以用更清晰的形式来编写长的 if-then-else 链
func switch2()  {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	switch1()
	switch2()
}