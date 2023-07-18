package test_tool

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, " : ", i)
	}
}
func Goroutinetest() {
	fmt.Println("go routine start:")

	f("direct f")

	go f("go routine f")

	go func(msg string) {
		fmt.Println("concurrently go func:", msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
	fmt.Println("go routine end")
}
