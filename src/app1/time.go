package main

import (
	"fmt"
	"time"
)

func main() {
	base()
}

func base() {
	timeNow := time.Now() //2021-11-11 15:15:27.839128695 +0800 CST m=+0.000066795
	timeNowStr := timeNow.String()
	timeNowUnix := timeNow.Unix()
	timeNowUnixNano := timeNow.UnixNano()
	//y, m, d := timeNow.Date()

	year, mouth, day, hour, minute, second := timeNow.Year(), timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second()

	timeDate := time.Date(year, mouth, day, hour, minute, second, 0, time.Local)

	fmt.Println("timeNow:", timeNow)
	fmt.Println("timeNowStr:", timeNowStr)
	fmt.Println("timeNowUnix:", timeNowUnix)
	fmt.Println("timeNowUnixNano:", timeNowUnixNano)
	fmt.Println("timeDate:", timeDate)
	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, mouth, day, hour, minute, second)
}
