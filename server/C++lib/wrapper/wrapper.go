package wrapper
/*
#include "../lib/wrapper.hpp"
#cgo LDFLAGS: -L. -lwrapper
*/
import "C"

import (
	"unsafe"
)

func GetVectorDataSum(data []int32,dataLen int ) int{
/*	for i:=0;i<dataLen;i++{
		fmt.Printf("go the data is:%v \r\n",data[i])
		fmt.Printf("go the data addr is:0x%x\r\n",&data[i])
	}*/

	cPointer := (*C.int)((unsafe.Pointer)(&data[0]))
	sum := C.GetVectorDataSum(cPointer,C.int(dataLen))
	return int(sum)
}