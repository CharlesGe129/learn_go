package main

import "fmt"

type ObjType uint32

func main() {
	a := ObjType(25)
	fmt.Println(a < ObjType(25))
	fmt.Println(a == ObjType(25))
	fmt.Println(a > ObjType(25))
}
