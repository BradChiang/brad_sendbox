package test_tool

import (
	"fmt"
	"strconv"
)

func Numparse() {
	fmt.Println("num parse test start:")
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	s, _ := strconv.Atoi("135")
	fmt.Println(s)
	_, e := strconv.Atoi("test")
	fmt.Println(e)
	fmt.Println("num parse test end")
}
