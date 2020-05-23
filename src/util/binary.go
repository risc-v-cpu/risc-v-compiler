package util

import (
	"strconv"
	"strings"
)

func IntToBinaryString(num int, zeroFill int) string {
	s := ""

	for ; num > 0; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}

	needFillCount := zeroFill - len(s)
	if needFillCount > 0 {
		s = strings.Repeat("0", needFillCount) + s
	}

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
