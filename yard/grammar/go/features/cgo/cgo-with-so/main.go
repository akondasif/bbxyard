package main

/*
#cgo CFLAGS : -I./include
#cgo LDFLAGS: -L./lib -lyh_cgo_test
#include "test.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	cs := C.test_hello(C.CString("C-IAmC-UseSo"))
	fmt.Println(C.GoString(cs))
	C.free(unsafe.Pointer(cs))

	cs2 := C.test_hello(C.CString("C-IAmC-UseSo-2"))
	fmt.Println("Another:", C.GoString(cs2))
	C.yh_c_buffer_free(unsafe.Pointer(cs2))

	fmt.Println("done")
}
