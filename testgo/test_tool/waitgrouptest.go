package test_tool

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Println("Worker ", id, "starting")

	time.Sleep(time.Second)
	fmt.Println("Worker ", id, "done")
}

func Waitgrouptest() {

	fmt.Println("waitgroup test start:")
	var wg sync.WaitGroup //wait all goroutine finish

	for i := 1; i <= 5; i++ {
		wg.Add(1) //increase waitgroup counter

		id := i //avoid reused i in goroutine func

		go func() {
			defer wg.Done() //defer延遲執行done,直到go func結束才執行
			worker(id)
		}()
	}

	wg.Wait() //等待所有goroutine done

	fmt.Println("waitgroup test end")
}
