package wrapper

/*
#include "wrapper.h"
#include <stdlib.h> //for free
#cgo LDFLAGS: -L. -llib
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export F
func F(i C.int) {
	fmt.Printf("Go F(%d)\n", i)
}

func CallF5WithF() {
	C.call_f5_with_F()
}

func CallOtherFunctions() {
	// Simple C function call.
	C.f0()

	// Call with constant.
	C.f1(5)

	// Call with Go int - cast to C int.
	var i int = 8
	C.f1(C.int(i))

	// Convert Go string to C string (char*) and back.
	cs := C.CString("Go string")
	csRet := C.f2(cs)
	fmt.Printf("fmt: %s\n", C.GoString(csRet))
	C.free(unsafe.Pointer(cs))          // free memory right now ...
	defer C.free(unsafe.Pointer(csRet)) // ... or later

	// Pass C struct to C function by value.
	s1 := C.struct_s1{a: 5} // or cast with C.int: s1 := C.struct_s1{C.int(i)}
	s1Ret := C.f31(s1)
	fmt.Printf("f31: s1=%v, s1Ret=%v\n", s1, s1Ret)

	// Pass C struct to C function by pointer.
	s1 = C.struct_s1{a: 5}
	s1Ret = *C.f32(&s1)
	fmt.Printf("f32: s1=%v, s1Ret=%v\n", s1, s1Ret)

	// Pass C struct with int pointer to C function.
	// s2 := C.struct_s2{&C.int(i)} - compile error "cannot take the address of _Ctype_int(i)", create variable as below
	ci := C.int(i)
	s2 := C.struct_s2{p: &ci}
	fmt.Printf("s2=%v, *s2.p=%d, i=%d\n", s2, *s2.p, i)
	C.f4(s2)
	fmt.Printf("f4: s2=%v, *s2.p=%d, i=%d\n", s2, *s2.p, i)

	// Pass C function pointer to C function.
	// C.f5(C.f1)                 - compile error "must call C.f1"
	// fp := unsafe.Pointer(C.f1) - compile error "must call C.f1"
	// It is not possible to pass C function pointer via Go.

	// Pass C function pointer to C function with simple C wrapper (see top of this file).
	//C.call_f5_with_f1()
}

func AddFunc(a, b int) int {
	sum := C.f6(C.int(a), C.int(b))
	return int(sum)
}
