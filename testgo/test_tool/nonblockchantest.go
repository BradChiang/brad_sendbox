package test_tool

import (
	"fmt"
)

func Nonblockchantest() {

	fmt.Println("non-blocking channel test start")
	messages := make(chan string)
	signal := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("receive : ", msg)
	default:
		fmt.Println("still not receive the message")
	}

	msg := "Hi"

	select {
	case messages <- msg:
		fmt.Println("send message : ", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("receive message : ", msg)
	case sig := <-signal:
		fmt.Println("receive signal : ", sig)
	default:
		fmt.Println("no activity")
	}

	fmt.Println("non-block channel test end")
}
