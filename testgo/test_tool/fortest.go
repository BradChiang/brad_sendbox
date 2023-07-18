package test_tool

import "fmt"

func Forlooptest() {

	fmt.Println("for loop test:")
	fmt.Println("normal:")
	for i := 0; i <= 9; i++ {
		fmt.Println("for loop count : ", i)
	}

	fmt.Println("seperate:")
	i := 0
	for i <= 3 {
		fmt.Println("for loop count : ", i)
		i++
	}

	fmt.Println("break infinite loop:")
	for {
		fmt.Println("for loop break")
		break
	}

	fmt.Println("continue loop:")
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("for odd loop count : ", i)
	}
	fmt.Println("for loop test end")
}
