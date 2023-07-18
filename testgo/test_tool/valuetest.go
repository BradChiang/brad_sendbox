package test_tool

import "fmt"

func Value() {
	s1, s2 := "Go", "lang"
	fmt.Println("string:")
	fmt.Println(s1 + s2)

	fmt.Println("int:")
	fmt.Println("1 + 1 = ", 1+1)

	fmt.Println("float32:")
	fmt.Println("7.0 / 3.0 = ", float32(7.0/3.0))

	fmt.Println("float64:")
	fmt.Println("7.0 / 3.0 = ", float64(7.0/3.0))

	fmt.Println("LOGIC:")
	fmt.Println("T & F = ", true && false)
	fmt.Println("T | F = ", true || false)
	fmt.Println("not T = ", !true)

}
