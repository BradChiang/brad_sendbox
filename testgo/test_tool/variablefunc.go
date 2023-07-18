package test_tool

import "fmt"

func sum(num ...int) {
	sum := 0
	for _, n := range num {
		sum += n
	}
	fmt.Println("number list = ", num, " total = ", sum)
}

func Varifunctest() {
	fmt.Println("variadic function start")

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5, 6}
	sum(nums...) //使array轉成multi_input
	fmt.Println("variadic function end")
}
