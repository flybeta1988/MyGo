package interview

import (
	"fmt"
	"sync"
)

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	//ch := make(chan interface{}) //解除注释看看！
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()

		for key, value := range set.s {
			ch <- key
			println("Iter:", key, value)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func TestChannel() {
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	v := <-th.Iter()
	fmt.Printf("%s %v", "ch", v)
}
