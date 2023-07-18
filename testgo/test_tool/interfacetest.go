package test_tool

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type recte struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r recte) area() float64 {
	return r.height * r.width
}

func (r recte) perim() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area = ", g.area())
	fmt.Println("perim = ", g.perim())
}

func Interfacetest() {
	fmt.Println("interface test start:")
	r := recte{3, 4}
	c := circle{5}

	measure(r)
	measure(c)
	fmt.Println("interface test end")
}
