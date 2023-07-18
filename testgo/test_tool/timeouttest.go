package test_tool

import (
	"fmt"
	"time"
)

func Timeouttest() {

	fmt.Println("timeout test start:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "complete:1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "complete:2"
	}()

	select {
	case res := <-ch1:
		fmt.Println("receive : ", res)
	case <-time.After(time.Second):
		fmt.Println("Timeout : 1")
	}
	select {
	case res := <-ch2:
		fmt.Println("receive : ", res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout : 2")
	}
	fmt.Println("timeout test end")
}
