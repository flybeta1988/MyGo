package main

import (
	"fmt"
	"time"
)

type Vector []float64

func main() {
	v := &Vector{1.1, 1.2}
	fmt.Println(v)
	//v.DoAll()
}

func (v Vector) Sleep(f float64) float64 {
	time.Sleep(1e9)
	fmt.Println("sleep ok!")
	return f
}

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Sleep(v[i])
	}
}

const NCPU = 4

func (v Vector) DoAll(u Vector)  {
	c := make(chan int, NCPU) //用于接收每个CPU的任务完成信号

	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i + 1)*len(v)/NCPU, u, c)
	}

	//等待所有CPU的任务完成
	for i := 0; i < NCPU; i ++ {
		<-c //获取到一个数据，表示一个CPU计算完成
	}

	fmt.Println("finished!")
}
