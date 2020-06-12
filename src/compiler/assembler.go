package compiler

import (
	"log"
	"strconv"
	"strings"
)

type machineCodeStructure struct {
	OpCode  string
	Operand []int
}

func assemblerInstructionToMachineCodeStructure(assemblerInstruction string) machineCodeStructure {
	mcs := new(machineCodeStructure)

	opCodeEndIndex := strings.Index(assemblerInstruction, " ")
	operandStartIndex := opCodeEndIndex + 1

	opCode := strings.ToUpper(assemblerInstruction[:opCodeEndIndex])
	operandsWithComma := assemblerInstruction[operandStartIndex:]

	splitParts := strings.SplitN(operandsWithComma, ",", MaxOperand)

	operandIntArray := make([]int, 0)

	for _, v := range splitParts {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Covert operand string to int error. %v Please check the operand %v .", err, v)
		}
		operandIntArray = append(operandIntArray, i)
	}

	mcs.OpCode = opCode
	mcs.Operand = operandIntArray

	return *mcs
}
