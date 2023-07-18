package test_tool

import (
	"fmt"
	"math"
)

const st string = "ConstTest " + "start:"

func ConstTest() {

	fmt.Println(st)

	const n = 50000 * 10000

	const d = 3e20 / n

	fmt.Println("int 3e20 / 50000 * 10000 = ", d)
	fmt.Println("int64 3e20 / 50000 * 10000 = ", int64(d))
	fmt.Println("Sin(50000 * 10000) = ", math.Sin(n))

	s := "ConstTest " + "end"
	fmt.Println(s)

}
