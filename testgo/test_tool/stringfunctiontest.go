package test_tool

import (
	"fmt"
	s "strings"
)

var pln = fmt.Println
var pf = fmt.Printf

type point struct {
	x, y int
}

func Stringfunctest() {
	pln("string function test start:")
	pln("test contains es     : ", s.Contains("test", "es"))
	pln("test count t         : ", s.Count("test", "t"))
	pln("test hasprefix te    : ", s.HasPrefix("test", "te"))
	pln("test hassuffix st    : ", s.HasSuffix("test", "es"))
	pln("test index e         : ", s.Index("test", "e"))
	pln("join  {a,b} -        : ", s.Join([]string{"a", "b"}, "-"))
	pln("repeat a 5           : ", s.Repeat("a", 5))
	pln("foo replace o to 0   : ", s.Replace("foo", "o", "0", -1))
	pln("foo replace 1 o to 0 : ", s.Replace("foo", "o", "0", 1))
	pln("split -  a-b-c-d-e   : ", s.Split("a-b-c-d-e", "-"))
	pln("tolower TEST         : ", s.ToLower("TEST"))
	pln("toupper test         : ", s.ToUpper("test"))
	pln("string function test end")
}

func Stringformat() {
	pln("string format test start:")

	p := point{1, 2}
	pf("struct1 : %v\n", p)
	pf("struct2 : %+v\n", p)
	pf("struct3 : %#v\n", p)

	pf("type : %T\n", p)

	pf("true bool : %t\n", true)

	pf("int : %d\n", 123)
	pf("bin : %b\n", 14)
	pf("char : %c\n", 33)
	pf("hex : %x\n", 456)
	pf("float1 : %f\n", 78.9)
	pf("float2 : %e\n", 123400000.0)
	pf("float3 : %E\n", 123400000.0)

	pf("string1 : %s\n", "\"string\"")
	pf("string2 : %q\n", "\"string\"")
	pf("string3 : %x\n", "hex this")

	pf("pointer : %p\n", &p)

	pf("width1 : %6d|%6d|\n", 12, 345)
	pf("width2 : %6.2f|%6.2f|\n", 1.2, 3.45)
	pf("width3 : %-6.2f|%-6.2f|\n", 1.2, 3.45)
	pf("width4 : %6s|%6s|\n", "foo", "b")
	pf("width5 : %-6s|%-6s|\n", "foo", "b")

	tmp := fmt.Sprintf("sprintf: a %s", "string")
	pln(tmp)

	pln("string format test end")
}
