package test_tool

import (
	"fmt"
	"sort"
)

type bylen []string

func Sortingtest() {
	fmt.Println("sort test start:")
	s := []string{"c", "b", "a", "x", "t"}
	fmt.Println("String unsorted : ", s)
	sort.Strings(s)
	fmt.Println("String after sorted : ", s)

	i := []int{23, 5, 23, 675, 2, 77}
	fmt.Println("number unsorted : ", i)
	sort.Ints(i)
	fmt.Println("number after sorted : ", i)

	check := sort.IntsAreSorted(i)
	fmt.Println("check number is sorted : ", check)

	fmt.Println("sort test end")
}

func (s bylen) Len() int {
	return len(s)
}

func (s bylen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s bylen) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
func Sortfunctiontest() {
	fmt.Println("sort by function test start")

	s := []string{"apple", "banana", "pie", "orange"}

	fmt.Println("name unsorted : ", s)
	sort.Sort(bylen(s))
	fmt.Println("name sorted : ", s)

	fmt.Println("sort by function test end")
}
