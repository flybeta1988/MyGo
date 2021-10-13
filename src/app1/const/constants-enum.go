package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)

func main() {
	fmt.Println("foo:", Tuesday)
}
