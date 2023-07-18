package test_tool

import (
	"fmt"
	"time"
)

// 用以傳遞message給goroutine function, channel 在 insert data or read data時會block func
func Channeltest() {
	fmt.Println("channel test start")
	messages := make(chan string)

	go func() {
		messages <- "ping"
		close(messages)
	}()

	msg := <-messages
	fmt.Println(msg)

	fmt.Println("channel test end")
}
func Channelbuffertest() {
	fmt.Println("channel buffering test start")
	messages := make(chan string, 2)

	messages <- "Buffering"

	messages <- "Channel"

	msg := <-messages
	fmt.Println(msg)
	fmt.Println(<-messages)

	close(messages)

	fmt.Println("channel buffering test end")
}

func work(done chan bool) {
	fmt.Println("Start Working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ChannelSynctest() {

	fmt.Println("channel sync. test start:")
	done := make(chan bool, 1)
	go work(done)

	<-done
	fmt.Println("channel sync. test end")
}

func ch1(chan1 chan<- string, msg string) {
	chan1 <- msg
}
func ch2(chan1 <-chan string, chan2 chan<- string) {
	msg := <-chan1
	chan2 <- msg
}

// channel 在 func parameter中, 可以指定channel 只能進行 send or receive
func Channeldirectiontest() {
	fmt.Println("channel direction test start:")
	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)
	ch1(chan1, "passing msg")
	ch2(chan1, chan2)
	fmt.Println("receive :", <-chan2)
	fmt.Println("channel direction test end")
}

func Closingchanneltest() {

	fmt.Println("closing channel test start:")
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs //moreget false when jobs close
			if more {
				fmt.Println("receive job : ", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 0; j < 3; j++ {
		jobs <- j
		fmt.Println("send job : ", j)
	}
	close(jobs)
	fmt.Println("send all jobs")

	fmt.Println("closing channel test end")

}

func Rangeoverchanneltest() {
	fmt.Println("range over channel test start:")
	ch := make(chan string, 5)
	ch <- "brian"
	ch <- "brad"
	ch <- "lucy"
	close(ch)

	for elem := range ch {
		fmt.Println(elem)
	}

	fmt.Println("range over channel test end")
}
