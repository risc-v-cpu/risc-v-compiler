package generator

import (
	"fmt"
	"log"
)

func GenerateHex(input []byte) []byte {
	log.Println(len(input), input)

	result := ""
	format := "%s:%X%X%X%X%X\n"
	for address := 0; address < len(input)/4; address++ {
		inputBytes32bit := []byte{
			input[address*4+0],
			input[address*4+1],
			input[address*4+2],
			input[address*4+3],
		}
		hexLL := []byte{byte(len(inputBytes32bit))}
		hexAAAA := []byte{byte(address >> 16*16), byte(address)}
		hexTT := []byte{0x00}
		hexDD := inputBytes32bit

		hexCC := getHexCheck(hexLL, hexAAAA, hexTT, hexDD)
		result = fmt.Sprintf(format, result, hexLL, hexAAAA, hexTT, hexDD, hexCC)
	}

	// add hex file format EOF
	//hexBOF := ":040000030000"
	hexEOF := ":00000001FF"
	result = result + hexEOF

	//log.Println("result", result)

	byteArray := []byte(result)

	return byteArray
}

func getHexCheck(byteArray ...[]byte) uint8 {
	//log.Println(byteArray)
	sum := 0

	for _, item := range byteArray {
		for _, itemByte := range item {
			sum = sum + int(itemByte)
		}
	}

	sumLSB := uint8(sum)
	result := 0x100 - uint16(sumLSB)

	//log.Println(uint8(result))
	return uint8(result)
}
