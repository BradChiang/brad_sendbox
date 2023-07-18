package test_tool

import "fmt"

func get() (int, int, string) {
	return 112, 7, "brad"
}

func Multireturntest() {

	fmt.Println("Multi-return start")

	year, month, name := get()

	fmt.Println("Year = ", year, " | Month = ", month, " | Name = ", name)

	a, b, _ := get()

	fmt.Println("a = ", a, " | b = ", b)

	fmt.Println("Multi-return end")
}
