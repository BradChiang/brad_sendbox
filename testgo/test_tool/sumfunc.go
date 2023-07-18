package test_tool

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plus3(a, b, c int) int {
	return a + b + c
}

func Functest() {
	fmt.Println("function test start")

	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plus3(3, 4, 5)
	fmt.Println("3 + 4 + 5 = ", res)

	fmt.Println("function test end")

}
