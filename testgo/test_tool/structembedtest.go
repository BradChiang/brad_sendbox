package test_tool

import "fmt"

type base struct {
	num int
}

func (b base) describ() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

type container struct {
	base
	str string
}

func Embaddingtest() {
	fmt.Println("struct embedding test start:")

	co := container{
		base: base{num: 1},
		str:  "some name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("base num = ", co.base.num)
	fmt.Println("describe = ", co.describ())

	type describer interface {
		describ() string
	}
	var d describer = co
	fmt.Println("describer = ", d.describ())
	fmt.Println("struct embedding test end")
}
