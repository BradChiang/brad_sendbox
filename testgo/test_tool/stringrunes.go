package test_tool

import (
	"fmt"
	"unicode/utf8"
)

func examRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
func StringRunestest() {

	fmt.Println("String & Runes test start")

	const s = "สวัสดี"

	fmt.Println(s, " length : ", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Runes count : ", utf8.RuneCountInString(s)) //size of the string

	for id, runeval := range s {
		fmt.Printf("%#U starts at %d\n", runeval, id)
	}
	fmt.Println("decode runes in string:")
	for i, w := 0, 0; i < len(s); i += w {
		runeval, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeval, i)
		w = width
		examRune(runeval)
	}

	fmt.Println("String & Runes test end")
}
