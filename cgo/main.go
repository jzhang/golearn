package main

/*
#include <stdio.h>
#include <stdlib.h>

typedef int (*intFunc) ();
int bridge_int_func(intFunc f) {
	return f();
}

int fortytwo() {
	return 42;
}

static void myPrint(const char* msg) {
	printf("%s", msg);
}

extern void myprint(int i);

void dofoo(void) {
	int i;
	for (i=0;i<10;i++) {
		myprint(i);
	}
}

int cArray[] = {0, 2, 3, 4, 5, 6, 7, 8};

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func CArrayToGoArray(cArray unsafe.Pointer, size int) (goArray []int32) {
	p := uintptr(cArray)
	for i := 0; i < size; i++ {
		j := *(*int32)(unsafe.Pointer(p))
		goArray = append(goArray, j)
		p += C.sizeof_int
	}
	return
}

func main() {
	C.puts(C.CString("Hello, 世界\n"))
	C.myPrint(C.CString("Hello, 世界\n"))
	cStr := C.CString("Hello, 世界\n")
	defer C.free(unsafe.Pointer(cStr))
	f := C.intFunc(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))
	C.dofoo()
	fmt.Println(C.sizeof_int)
	goArray := CArrayToGoArray(unsafe.Pointer(&C.cArray[0]), 8)
	fmt.Println(goArray)

}
