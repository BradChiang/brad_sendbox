package test_tool

import "fmt"

func SingleArr() {

	fmt.Println("single array start")

	var a [5]int
	fmt.Println("empty : ", a)

	a[4] = 44
	fmt.Println("add item : ", a)
	fmt.Println("get item : ", a[4])
	fmt.Println("get length : ", len(a))

	b := []int{1, 2, 3, 4, 5}
	fmt.Println("Array b = ", b)

}

func TwoDimensionArr() {

	fmt.Println("two dimension array start")

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("two dimension array = ", twoD)
}
