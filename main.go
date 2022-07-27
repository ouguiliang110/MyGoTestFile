package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "慕课网moody"
	fmt.Println(len(s))
	fmt.Println(unsafe.Sizeof("慕"))
	fmt.Println(unsafe.Sizeof("慕课"))
	fmt.Println(unsafe.Sizeof("慕课网"))
	fmt.Println(unsafe.Sizeof("慕课网Moody"))
}
