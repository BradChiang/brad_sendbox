package test_tool

import "fmt"

func Rangetest() {
	//for 的 range 會抓出 array/slice/map 的 index_key, index_value, 不用設定其他條件式

	fmt.Println("range test start")
	num := []int{2, 4, 6, 8, 10}
	sum := 0
	for _, n := range num {
		sum += n
	}
	fmt.Println("sum = ", sum)

	for i, n := range num {
		if n%3 == 0 {
			fmt.Println("Index ", i, " value is devisible by 3 : ", n)
		}
	}

	m := map[string]string{"brian": "home", "brad": "company"}
	for k, v := range m {
		fmt.Println("map get : ", k, " -> ", v)
	}
	fmt.Println("just get map key")
	for k := range m {
		fmt.Println("key = ", k)
	}
	fmt.Println("get string char")
	for i, c := range "Hallo World" {
		fmt.Println("NO.", i, " Char byte = ", c, "| Char = ", string(c))
	}
	fmt.Println("range test end")
}
