package test_tool

import (
	"fmt"
	"time"
)

func Timertest() {
	fmt.Println("timer test start:")
	timer1 := time.NewTimer(2 * time.Second)

	fmt.Println("Working...")
	<-timer1.C
	fmt.Println("Timer 1 alart : Timeout")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 alart : Timeout")
	}()

	done := timer2.Stop()

	if done {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
	fmt.Println("timer test end")
}

// tickers = repeatledli timer
func Tickertest() {

	fmt.Println("ticket test start")
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at : ", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
	fmt.Println("ticket test end")
}
