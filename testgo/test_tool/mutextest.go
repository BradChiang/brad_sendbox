package test_tool

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func Mutextest() {
	fmt.Println("mutex test start:")

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrease := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrease("a", 10000)
	go doIncrease("a", 10000)
	go doIncrease("b", 10000)

	wg.Wait()

	fmt.Println(c.counters)
	fmt.Println("mutex test end")
}
