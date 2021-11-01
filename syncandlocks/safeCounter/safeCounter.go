package main

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	i int
	sync.Mutex
}

func main() {
	sc := new(safeCounter)

	for i := 0; i < 100; i++ {
		go sc.Increment()
		go sc.Decrement()
	}

	fmt.Println(sc.GetValue())
}

func (sc *safeCounter) Increment() {
	sc.Lock()
	sc.i++
	sc.Unlock()
}

func (sc *safeCounter) Decrement() {
	sc.Lock()
	sc.i--
	sc.Unlock()
}

func (sc *safeCounter) GetValue() int {
	
	var result int
	
	sc.Lock()
	result = sc.i
	sc.Unlock()
	
	return result;
}
