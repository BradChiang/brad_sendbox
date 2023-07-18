package test_tool

import (
	"fmt"
	"time"
)

func Switchtest(i int) {

	fmt.Println("switch test start:")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("munber:", i)
	}

	fmt.Println("Time switch:")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is weekend : ", time.Now().Weekday())
	default:
		fmt.Println("Today is weekday : ", time.Now().Weekday())
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Time is before moon")
	default:
		fmt.Println("Time is after moon")
	}

	TypeSwitvh := func(x interface{}) {
		switch tmp := x.(type) {
		case bool:
			fmt.Println("Bool type")
		case int, float32, float64:
			fmt.Println("number type")
		case string:
			fmt.Println("string type")
		default:
			fmt.Printf("unknown type : %T\n", tmp)
		}
	}
	var a []byte
	TypeSwitvh(true)
	TypeSwitvh(23456)
	TypeSwitvh("hallo world")
	TypeSwitvh(a)

	fmt.Println("switch case end")
}
