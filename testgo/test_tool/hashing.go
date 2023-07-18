package test_tool

import (
	"crypto/sha256"
	"fmt"
)

func Hashing() {
	fmt.Println("sha256 hash test start")
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
	fmt.Println("sha256 hash test end")
}
