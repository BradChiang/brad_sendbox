package test_tool

import "fmt"

var start = "variable output:"

func Vartest() {

	fmt.Println(start)
	var oi int
	fmt.Println("without initialization int = ", oi)
	var of3 float32
	fmt.Println("without initialization float32 = ", of3)
	var of6 float64
	fmt.Println("without initialization float64 = ", of6)
	var os string
	fmt.Println("without initialization string = ", os)
	var ob bool
	fmt.Println("without initialization bool = ", ob)

	var a = "initial"
	fmt.Println("var a : ", a)
	var b, c int = 5, 10
	fmt.Println("var b,c int : ", b, c)
	var e float32 = 2
	fmt.Println("var e float32 : ", e)
	var f float64 = 2
	fmt.Println("var f float64 : ", f)
	var g = true
	fmt.Println("var g : ", g)
	s := `end of 
		var test`
	fmt.Println(s)
}
