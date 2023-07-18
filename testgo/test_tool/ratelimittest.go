package test_tool

import (
	"fmt"
	"time"
)

func Ratelimittest() {

	fmt.Println("rate limit test start")

	reqs := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		reqs <- i
	}
	close(reqs)
	limiter := time.NewTicker(200 * time.Millisecond)

	for req := range reqs {
		<-limiter.C
		fmt.Println("request : ", req, time.Now())
	}

	burstylimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstylimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstylimiter <- t
		}
	}()

	burstyreqs := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyreqs <- i
	}
	close(burstyreqs)

	for req := range burstyreqs {
		<-burstylimiter
		fmt.Println("request : ", req, time.Now())
	}
	limiter.Stop()
	fmt.Println("rate limit test end")
}
