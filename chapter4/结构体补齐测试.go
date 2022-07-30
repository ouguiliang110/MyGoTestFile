package main

import (
	"fmt"
	"unsafe"
)

type K struct {
	a int64
	s string
	b struct{}
}

func main() {
	fmt.Println(unsafe.Sizeof(K{}))
}
