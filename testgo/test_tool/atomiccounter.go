package test_tool

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCountertest() {
	fmt.Println("atomic counter test start:")

	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops = ", ops)
	fmt.Println("atomic counter test end")
}
