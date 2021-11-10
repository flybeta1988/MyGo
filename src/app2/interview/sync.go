package interview

import (
	"fmt"
	"sync"
)

//考点：map线程安全

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func TestLock() {
	ua := &UserAges{}

	go func(ua *UserAges, name string) {
		ua.Add(name, 30)
		fmt.Println("add ok")
	}(ua, "a")

	go func(ua *UserAges, name string) {
		age := ua.Get(name)
		fmt.Println("get ok, age:", age)
	}(ua, "a")

	fmt.Println(ua.ages)
}
