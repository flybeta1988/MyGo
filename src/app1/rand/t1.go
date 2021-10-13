package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rInt := rand.Int()
	rInt31 := rand.Int31()
	rInt31n := rand.Int31n(10)
	rInt63 := rand.Int63()
	rInt63n := rand.Int63n(10)

	r := mtrand()

	fmt.Println("int:", rInt)
	fmt.Println(rInt31)
	fmt.Println(rInt31n)
	fmt.Println(rInt63)
	fmt.Println(rInt63n)
	fmt.Println(r)
}

func mtrand() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
