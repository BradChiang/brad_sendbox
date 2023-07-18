package main

import (
	"fmt"
	"os"
)

// using go build & use command line ./commandlineargu a b c d
func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

}
