package compiler

import "log"

const (
	MaxOperand = 4
)

type rv32MachineCode [4]byte

func CompileAll(assemblerInstructions []string) []byte {
	result := make([]byte, 0)
	for _, v := range assemblerInstructions {
		byteArray := compile(v)
		for _, v1 := range byteArray {
			result = append(result, v1)
		}
	}
	return result
}

func compile(assemblerInstruction string) []byte {
	result := make([]byte, 0)
	mcs := assemblerInstructionToMachineCodeStructure(assemblerInstruction)
	byteArray32bit := mcs.machineCodeStructureToByteArray()
	log.Printf("%v opcode: %v funct3: %v %v \t  \t %v \n", byteArray32bit, byteArray32bit[31-6:31-0+1], byteArray32bit[31-14:31-12+1], mcs, assemblerInstruction)
	byteArray4bit := convert32ByteArrayFrom32bitTo4bit(byteArray32bit)
	for _, v := range byteArray4bit {
		result = append(result, v)
	}
	return result
}

func convert32ByteArrayFrom32bitTo4bit(byteArray32bit rv32) rv32MachineCode {
	var machineCode rv32MachineCode
	for k, v := range byteArray32bit {
		byteIndex := k / 8
		bit := 7 - k % 8
		machineCode[byteIndex] = machineCode[byteIndex] | v << bit
	}
	return machineCode
}

func (mcs machineCodeStructure)machineCodeStructureToByteArray() rv32 {
	var result rv32

	rv32Type := RV32_TYPE_NULL

	switch mcs.OpCode {

	case "BEQ"	,"BNE"	,"BLT"	,"BGE"	,"BLTU"	,"BGEU"	:
		rv32Type = RV32_TYPE_B

	case "SB"	,"SH"	,"SW"	:
		rv32Type = RV32_TYPE_S

	case "LB"	,"LH"	,"LW"	,"LBU"	,"LHU"	:
		rv32Type = RV32_TYPE_L

	case "ADDI"	,"SLTI"	,"SLTIU"	,"XORI"	,"ORI"	,"ANDI"	,"SLLI"	,"SRLI"	,"SRAL"	:
		rv32Type = RV32_TYPE_I

	case "ADD"	,"SUB"	,"SLL"	,"SLT"	,"SLTU"	,"XOR"	,"SRL"	,"SRA"	,"OR"	,"AND"	:
		rv32Type = RV32_TYPE_R

	}

	switch rv32Type {
	case RV32_TYPE_I:
		result.setOpCode(RV32_OPCODE_I)
		result.setRd(mcs.Operand[0])
		result.setRs1(mcs.Operand[1])
		result.setImmBitForRv32TypeI(mcs.Operand[2])

	case RV32_TYPE_S:
		result.setOpCode(RV32_OPCODE_S)
		result.setRs1(mcs.Operand[0])
		result.setRs2(mcs.Operand[1])
		result.setImmBitForRv32TypeS(mcs.Operand[2])

	case RV32_TYPE_R:
		result.setOpCode(RV32_OPCODE_R)
		result.setRd(mcs.Operand[0])
		result.setRs1(mcs.Operand[1])
		result.setRs2(mcs.Operand[2])

	case RV32_TYPE_L:
		result.setOpCode(RV32_OPCODE_L)
		result.setRd(mcs.Operand[0])
		result.setRs1(mcs.Operand[1])
		result.setImmBitForRv32TypeI(mcs.Operand[2])

	case RV32_TYPE_B:
		result.setOpCode(RV32_OPCODE_B)
		result.setRs1(mcs.Operand[0])
		result.setRs2(mcs.Operand[1])
		result.setImmBitForRv32TypeB(mcs.Operand[2])

	}

	switch mcs.OpCode {

	case "BEQ"	:
		result.setFunct3(RV32_FUNCT3_BEQ)
	case "BNE"	:
		result.setFunct3(RV32_FUNCT3_BNE)
	case "BLT"	:
		result.setFunct3(RV32_FUNCT3_BLT)
	case "BGE"	:
		result.setFunct3(RV32_FUNCT3_BGE)
	case "BLTU"	:
		result.setFunct3(RV32_FUNCT3_BLTU)
	case "BGEU"	:
		result.setFunct3(RV32_FUNCT3_BGEU)

	case "SB"	:
		result.setFunct3(RV32_FUNCT3_SB)
	case "SH"	:
		result.setFunct3(RV32_FUNCT3_SH)
	case "SW"	:
		result.setFunct3(RV32_FUNCT3_SW)

	case "LB"	:
		result.setFunct3(RV32_FUNCT3_LB)
	case "LH"	:
		result.setFunct3(RV32_FUNCT3_LH)
	case "LW"	:
		result.setFunct3(RV32_FUNCT3_LW)
	case "LBU"	:
		result.setFunct3(RV32_FUNCT3_LBU)
	case "LHU"	:
		result.setFunct3(RV32_FUNCT3_LHU)

	case "ADDI"	:
		result.setFunct3(RV32_FUNCT3_ADDI)
	case "SLTI"	:
		result.setFunct3(RV32_FUNCT3_SLTI)
	case "SLTIU"	:
		result.setFunct3(RV32_FUNCT3_SLTIU)
	case "XORI"	:
		result.setFunct3(RV32_FUNCT3_XORI)
	case "ORI"	:
		result.setFunct3(RV32_FUNCT3_ORI)
	case "ANDI"	:
		result.setFunct3(RV32_FUNCT3_ANDI)
	case "SLLI"	:
		result.setFunct3(RV32_FUNCT3_SLLI)
		result.setFunct7(RV32_FUNCT7_0000000)
		result.setRs2(mcs.Operand[2])
	case "SRLI"	:
		result.setFunct3(RV32_FUNCT3_SRLI)
		result.setFunct7(RV32_FUNCT7_0000000)
		result.setRs2(mcs.Operand[2])
	case "SRAL"	:
		result.setFunct3(RV32_FUNCT3_SRAL)
		result.setFunct7(RV32_FUNCT7_0100000)
		result.setRs2(mcs.Operand[2])

	case "ADD"	:
		result.setFunct3(RV32_FUNCT3_ADD)
	case "SUB"	:
		result.setFunct3(RV32_FUNCT3_SUB)
		result.setFunct7(RV32_FUNCT7_0100000)
	case "SLL"	:
		result.setFunct3(RV32_FUNCT3_SLL)
	case "SLT"	:
		result.setFunct3(RV32_FUNCT3_SLT)
	case "SLTU"	:
		result.setFunct3(RV32_FUNCT3_SLTU)
	case "XOR"	:
		result.setFunct3(RV32_FUNCT3_XOR)
	case "SRL"	:
		result.setFunct3(RV32_FUNCT3_SRL)
	case "SRA"	:
		result.setFunct3(RV32_FUNCT3_SRA)
		result.setFunct7(RV32_FUNCT7_0100000)
	case "OR"	:
		result.setFunct3(RV32_FUNCT3_OR)
	case "AND"	:
		result.setFunct3(RV32_FUNCT3_AND)
	case "JAL"	:
		result.setOpCode(RV32_OPCODE_JAL)
		result.setImmBitForRv32TypeJ(mcs.Operand[0])
		result.setRd(mcs.Operand[1])
	case "JALR"	:
		result.setOpCode(RV32_OPCODE_JALR)
	case "AUIPC":
		result.setOpCode(RV32_OPCODE_AUIPC)
	case "LUI"	:
		result.setOpCode(RV32_OPCODE_LUI)

	}


	return result
}
