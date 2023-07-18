package test_tool

import (
	"encoding/xml"
	"fmt"
)

type plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p plant) toString() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func Xmltest() {
	fmt.Println("xml test start:")
	coffee := &plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	out, _ := xml.MarshalIndent(coffee, " ", "   ")
	fmt.Println(string(out))
	fmt.Println(xml.Header, string(out))

	var p plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p.toString())

	tomato := &plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", "   ")
	fmt.Println(string(out))

	fmt.Println("xml test end")
}
