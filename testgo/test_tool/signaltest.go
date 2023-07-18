package test_tool

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Signaltest() {
	fmt.Println("signal test start:")
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awating signal")
	<-done
	fmt.Println("exit")
	fmt.Println("signal test end")
}
