package test_tool

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { //strucct method
	return r.width * r.height
}

func (r rect) perim() int { //strucct method
	return 2*r.width + 2*r.height
}

func Methodtest() {
	fmt.Println("Method test start:")

	r := rect{width: 10, height: 5}
	fmt.Println("height = ", r.height, "width = ", r.width)
	fmt.Println("area = ", r.area())
	fmt.Println("perim = ", r.perim())

	rp := &r
	fmt.Println("height = ", rp.height, "width = ", rp.width)
	fmt.Println("area = ", rp.area())
	fmt.Println("perim = ", rp.perim())
	fmt.Println("Method test end")

}
