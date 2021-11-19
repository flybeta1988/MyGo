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
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func TestLock() {
	ua := &UserAges{ages: make(map[string]int)}

	var wg sync.WaitGroup
	wg.Add(2)

	go func(ua *UserAges, name string) {
		ua.Add(name, 30)
		fmt.Println("add ok")
		wg.Done()
	}(ua, "a")

	go func(ua *UserAges, name string) {
		age := ua.Get(name)
		fmt.Println("get ok, age:", age)
		wg.Done()
	}(ua, "a")

	wg.Wait()

	fmt.Println(ua.ages)
}
