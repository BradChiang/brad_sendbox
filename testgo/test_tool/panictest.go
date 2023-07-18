package test_tool

import (
	"fmt"
	"os"
)

// panic = unexpect wrong
func Panictest() {
	fmt.Println("Panic test start:")

	panic("a problem")

	_, err := os.Create("/tmp/file")

	if err != nil {
		panic(err)
	}
	fmt.Println("Panic test end")
}
