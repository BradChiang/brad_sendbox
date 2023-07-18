package test_tool

import (
	"embed"
	"fmt"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func Embaddir() {
	fmt.Println("embad directive test start")
	print(fileString)
	print(string(fileByte))

	cont1, _ := folder.ReadFile("folder/file1.hash")
	cont2, _ := folder.ReadFile("folder/file2.hash")
	print(string(cont1))
	print(string(cont2))
	fmt.Println("embad directive test end")
}
