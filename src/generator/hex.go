package generator

import (
	"fmt"
)

func GenerateHex(input []byte) []byte {
	result := ""

	hexLL := []byte{byte(len(input))}
	hexAAAA := []byte{0x00, 0x00}
	hexTT := []byte{0x00}
	hexDD := input
	hexCC := getHexCheck(hexLL, hexAAAA, hexTT, hexDD)

	// add hex file format EOF
	hexEOF := ":00000001FF"

	format := "%s:%x%x%x%x%x\n%s\n"
	result = fmt.Sprintf(format, result, hexLL, hexAAAA, hexTT, hexDD, hexCC, hexEOF)

	//log.Println("result", result)

	byteArray := []byte(result)

	return byteArray
}

func getHexCheck(byteArrayArray ...[]byte) uint8 {
	result := 0
	var sum = 0

	for _, byteArray := range byteArrayArray {
		for _, v := range byteArray {
			sum = sum + int(v)
		}
	}
	result = sum % 0x100
	return uint8(result)
}
