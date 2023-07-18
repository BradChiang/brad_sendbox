package test_tool

import "fmt"

func myp() {
	panic("a problem")
}

func Recovertest() {
	fmt.Println("Recover test start:")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error : ", r)
		}
	}()
	myp()
	fmt.Println("panic occur")
	fmt.Println("Recover test end")
}
