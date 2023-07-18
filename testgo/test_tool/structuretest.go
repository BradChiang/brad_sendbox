package test_tool

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name: name, age: age}
	return &p
}

func Structuretest() {
	fmt.Println("struct test start:")

	fmt.Println(person{"Brian", 23})
	fmt.Println(person{name: "Lucy", age: 26})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 55})
	fmt.Println(newPerson("John", 40))

	p := person{name: "Sean", age: 50}
	fmt.Println(p.name)

	sp := &p
	fmt.Println("get age : ", sp.age)

	sp.age = 10
	fmt.Println("change age : ", sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println("Dog info : ", dog)

	fmt.Println("struct test end")
}
