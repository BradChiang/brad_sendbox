package test_tool

import (
	"bytes"
	"fmt"
	"regexp"
)

func Regularexpression() {

	fmt.Println("regular expression test start")
	fmt.Println("regular expression = p([a-z]+)ch")
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("peach match:", match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	teststr := "peach punch"
	fmt.Println("string: ", teststr)

	fmt.Println("find regular string : ", r.FindString(teststr))

	fmt.Println("index : ", r.FindStringIndex(teststr))
	fmt.Println("find submatch : ", r.FindStringSubmatch(teststr))
	fmt.Println("find submatch index: ", r.FindStringSubmatchIndex(teststr))
	teststr = "peach punch pinch"
	fmt.Println("all regular string: ", r.FindAllString(teststr, -1))
	fmt.Println("all index : ", r.FindAllStringSubmatchIndex(teststr, -1))
	fmt.Println("find first 2 match regular string : ", r.FindAllString("peach punch pinch", 2))
	fmt.Println([]byte("peach"), " byte match :", r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp : ", r)
	fmt.Println("string : a peach")
	fmt.Println("replace regular string:", r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

	fmt.Println("regular expression test end")

}
