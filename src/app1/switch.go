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

func swithc3() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Printf("Today's type:%T\n", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

//类型选择
func switchType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I dont't know about type %T\n", v)
	}
}

func main() {
	//switch1()
	//switch2()
	//swithc3()
	switchType(1)
}