package test_tool

import "fmt"

func IfElse(num int) {

	fmt.Println("if-else test:")

	if num%2 == 0 {
		fmt.Println(num, " is even number")
	} else {
		fmt.Println(num, " is odd number")
	}

	if num%4 == 0 {
		fmt.Println(num, " is devisible by 4")
	} else {
		fmt.Println(num, " is not devisible by 4")
	}

	if i := num; i < 0 {
		fmt.Println(num, " is negative")
	} else if i < 10 {
		fmt.Println(num, " has one digit")
	} else {
		fmt.Println(num, " has multi-digits")
	}
	fmt.Println("if-else test end")
}
