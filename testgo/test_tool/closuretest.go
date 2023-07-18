package test_tool

import "fmt"

func intSeq() func() int { //func intSeq do the init & return the function func() noname
	i := 0
	return func() int { //func() noname return the int i++
		i++
		return i
	}
}

func Closuretest() {

	fmt.Println("Closure test start")

	nextInt := intSeq()
	fmt.Println("Closure nextInt = ", nextInt())
	fmt.Println("Closure nextInt = ", nextInt())
	fmt.Println("Closure nextInt = ", nextInt())

	newInt := intSeq()
	fmt.Println("Closure newInt = ", newInt())

	fmt.Println("Closure test end")
}
