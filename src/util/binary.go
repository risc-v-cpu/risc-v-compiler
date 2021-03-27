package util

import (
	"strconv"
	"strings"
	"unsafe"
)

func IntToBinaryString(num int, fill int) string {
	s := strconv.FormatUint(*(*uint64)(unsafe.Pointer(&num)), 2)
	//log.Println("num", num, "fill", fill, "s", s)

	needFillCount := fill - len(s)
	if needFillCount >= 0 {
		extendNum := "0"
		if num < 0 {
			extendNum = "1"
		}
		s = strings.Repeat(extendNum, needFillCount) + s
	} else {
		s = s[-needFillCount:]
	}

	//log.Println("num", num, "fill", fill, "s", s)

	return s
}

//func Int2Byte(data int) (ret []byte) {
//	var length uintptr = unsafe.Sizeof(data)
//	ret = make([]byte, length)
//	var tmp int = 0xff
//	var index uint = 0
//	for index=0; index<uint(length); index++{
//		ret[index] = byte((tmp<<(index*8) & data)>>(index*8))
//	}
//	return ret
//}
//
//func Byte2Int(data []byte) int {
//	var ret int = 0
//	var length int = len(data)
//	var i uint = 0
//	for i=0; i<uint(length); i++{
//		ret = ret | (int(data[i]) << (i*8))
//	}
//	return ret
//}
