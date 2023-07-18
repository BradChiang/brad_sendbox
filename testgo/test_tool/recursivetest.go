package test_tool

import "fmt"

func recursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursive(n-1)
}

func Recursiontest() {

	fmt.Println("Recursion test start")

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println("recursive 7 = ", recursive(7))
	fmt.Println("fib 7 = ", fib(7))
	fmt.Println("Recursion test end")
}
