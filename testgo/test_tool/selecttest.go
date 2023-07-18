package test_tool

import (
	"fmt"
	"time"
)

func Selecttest() {

	fmt.Println("select test start:")
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch3 <- "first complete:3"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "second complete:2"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch4 <- "third complete:4"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- "last complete:1"
	}()

	for i := 0; i < 4; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("receive : ", msg)
		case msg := <-ch2:
			fmt.Println("receive : ", msg)
		case msg := <-ch3:
			fmt.Println("receive : ", msg)
		case msg := <-ch4:
			fmt.Println("receive : ", msg)
		}
	}
	fmt.Println("select test end")
}
