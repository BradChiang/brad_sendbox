package test_tool

import (
	"fmt"
	"time"
)

func workers(wid int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", wid, " start job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker", wid, " finish job ", j)
		results <- j * 2
	}
}

func Workerpooltest() {

	fmt.Println("workerpools test start:")

	const numJob = 5
	jobs := make(chan int, numJob)
	results := make(chan int, numJob)

	for w := 1; w <= 3; w++ {
		go workers(w, jobs, results)
	}
	for j := 1; j <= numJob; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJob; a++ {
		<-results
	}

	fmt.Println("workerpools test end")
}
