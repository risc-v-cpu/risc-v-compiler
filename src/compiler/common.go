package compiler

import (
	"risc-v-compiler/util"
)

type rv32 [32]byte

func (raw *rv32) setOpCode(strCode string) {
	raw.setBitForRv32(strCode, 0)
}

func (raw *rv32) setRd(rd int) {
	strCode := util.IntToBinaryString(rd, 5)
	raw.setBitForRv32(strCode, 7)
}

func (raw *rv32) setFunct3(strCode string) {
	raw.setBitForRv32(strCode, 12)
}

func (raw *rv32) setRs1(rs1 int) {
	strCode := util.IntToBinaryString(rs1, 5)
	raw.setBitForRv32(strCode, 15)
}

func (raw *rv32) setRs2(rs2 int) {
	strCode := util.IntToBinaryString(rs2, 5)
	raw.setBitForRv32(strCode, 20)
}

func (raw *rv32) setFunct7(strCode string) {
	raw.setBitForRv32(strCode, 25)
}

func (raw *rv32) setBitForRv32(strCode string, startIndex int) {
	length := len(strCode)
	for i := 0; i < length; i++ {
		num := strCode[i] - '0'
		index := 32-startIndex-(length-i)
		raw[index] = num
	}
	//return raw
}

func (raw *rv32) setImmBitForRv32TypeI(imm int) {
	//strImm := strconv.Itoa(imm)
	//strImm_11_0 = strImm[0:11]
	strCode := util.IntToBinaryString(imm, 12)
	raw.setBitForRv32(strCode, 20)
}

func (raw *rv32) setImmBitForRv32TypeS(imm int) {
	strCode := util.IntToBinaryString(imm, 12)
	strCode = util.ReverseString(strCode)
	strCode_4_0 := util.ReverseString(strCode[0:4+1])
	strCode_11_5 := util.ReverseString(strCode[5:11+1])
	//strCode_4_0 := strCode[0:4+1]
	//strCode_11_5 := strCode[5:11+1]
	raw.setBitForRv32(strCode_4_0, 7)
	raw.setBitForRv32(strCode_11_5, 25)
	//return result
}

func (raw *rv32) setImmBitForRv32TypeB(imm int) {
	//imm = imm / 10
	//TODO:fix bug
	strCode := util.IntToBinaryString(imm, 12+1)
	strCode = util.ReverseString(strCode)
	strCode_4_1 := util.ReverseString(strCode[1:4+1])
	strCode_11 := string(strCode[11])
	strCode_10_5 := util.ReverseString(strCode[5:10+1])
	strCode_12 := string(strCode[12])
	raw.setBitForRv32(strCode_4_1, 8)
	raw.setBitForRv32(strCode_11, 7)
	raw.setBitForRv32(strCode_10_5, 25)
	raw.setBitForRv32(strCode_12, 31)
	//return result
}

func (raw *rv32) setImmBitForRv32TypeU(imm int) {
	strCode := util.IntToBinaryString(imm, 32)
	strCode = util.ReverseString(strCode)
	strCode_31_12 := util.ReverseString(strCode[12:31+1])
	raw.setBitForRv32(strCode_31_12, 12)
	//return result
}

func (raw *rv32) setImmBitForRv32TypeJ(imm int) {
	imm = imm >> 1
	strCode := util.IntToBinaryString(imm, 32)
	strCode = util.ReverseString(strCode)
	strCode_20 := util.ReverseString(strCode[20:20+1])
	raw.setBitForRv32(strCode_20, 31)
	strCode_10_1 := util.ReverseString(strCode[1:10+1])
	raw.setBitForRv32(strCode_10_1, 31-1)
	strCode_11 := util.ReverseString(strCode[11:11+1])
	raw.setBitForRv32(strCode_11, 31-1-11)
	strCode_19_12 := util.ReverseString(strCode[12:19+1])
	raw.setBitForRv32(strCode_19_12, 31-1-11-1)
	//return result
}
