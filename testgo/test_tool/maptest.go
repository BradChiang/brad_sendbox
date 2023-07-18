package test_tool

import "fmt"

func Maptest() {

	fmt.Println("map test start")

	m := make(map[string]int)

	m["k05"] = 25
	m["k19"] = 19

	fmt.Println("map = ", m)

	v1 := m["k05"]
	fmt.Println("v1 = k05 -> ", v1)
	v0 := m["k00"]
	fmt.Println("v0 = k00 -> ", v0)

	fmt.Println("map length = ", len(m))

	delete(m, "k19")
	fmt.Println("delete k19 from map = ", m)

	_, prs := m["k19"]
	fmt.Println("k19 present : ", prs)

	n := map[string]int{"brian": 4129889, "brad": 1129084}
	fmt.Println("map n = ", n)
}
