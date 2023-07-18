package test_tool

import "fmt"

func Slicetest() {

	fmt.Println("Slices test start")

	var s0 []string
	fmt.Println("uninit : ", s0, " | array_null : ", s0 == nil, " | length = ", len(s0))

	s0 = make([]string, 3)
	fmt.Println("empty : ", s0, " | length = ", len(s0), " | capacity : ", cap(s0))

	s0[0] = "HALLO"
	s0[1] = "WORLD"
	s0[2] = "!!!"
	fmt.Println("set string : ", s0)
	fmt.Println("get string[2] : ", s0[2])
	fmt.Println("get length : ", len(s0))

	s0 = append(s0, "GO")
	s0 = append(s0, "LANG", "TEST")
	fmt.Println("append string : ", s0)

	c := make([]string, len(s0))
	copy(c, s0)
	fmt.Println("copy string arr : ", c)

	l := s0[2:5]
	fmt.Println("slice 2~4 : ", l)

	l = s0[:5]
	fmt.Println("slice 0~4 : ", l)

	l = s0[2:]
	fmt.Println("slice 2~end : ", l)

	t := []string{"intrising", "com", "tw"}
	for i := 0; i < len(t); i++ {
		fmt.Print(t[i], "/")
	}
	fmt.Println()

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("two dimension arr = ", twoD)
	fmt.Println("slice test end")

}
