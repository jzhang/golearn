package main

import "fmt"
import "unsafe"
import "reflect"

func main() {
	var a []int
	b := []int{}

    fmt.Printf("a: %#v\n", (*reflect.SliceHeader) (unsafe.Pointer(&a)))
    fmt.Printf("b: %#v\n", (*reflect.SliceHeader) (unsafe.Pointer(&b)))
    fmt.Printf("a size: %d\n", unsafe.Sizeof(a))
}
