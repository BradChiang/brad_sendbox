package test_tool

import (
	"fmt"
	"os"
)

func createfile(s string) *os.File {
	fmt.Println("create file")
	f, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	return f
}

func writefile(f *os.File, msg string) {
	fmt.Println("writing file")
	fmt.Fprintln(f, msg)
}

func closefile(f *os.File) {
	fmt.Println("closong file")
	err := f.Close()
	if err != nil {
		fmt.Println("closed file error")
		fmt.Fprintf(os.Stderr, "error : %v\n", err)
		os.Exit(1)
	}
}

func Defertest() {
	fmt.Println("defer test start:")

	f := createfile("/tmp/test.txt")
	defer closefile(f)
	writefile(f, "hallo world")

	fmt.Println("defer test end")
}
