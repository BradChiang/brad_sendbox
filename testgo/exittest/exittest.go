package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("exit test start:")

	defer fmt.Println("exit test end")

	os.Exit(3)
}
