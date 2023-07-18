package test_tool

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func PointerTest() {

	fmt.Println("pointer test start")

	i := 1
	fmt.Println("initialize i = ", i, " pointer = ", &i)
	zeroval(i)
	fmt.Println("zeroval i = ", i, " pointer = ", &i)
	zeroptr(&i)
	fmt.Println("zeroptr i = ", i, " pointer = ", &i)

	fmt.Println("pointer test end")
}
