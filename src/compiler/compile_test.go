package compiler

import "testing"

type convert32ByteArrayFrom32bitTo4bitTestCase struct {
	input  rv32
	output rv32MachineCode
}

func checkRv32MachineCode(a, b rv32MachineCode) bool {
	for i := 0; i < 4; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestConvert32ByteArrayFrom32bitTo4bit(t *testing.T) {
	testCase := []convert32ByteArrayFrom32bitTo4bitTestCase{
		{rv32{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, rv32MachineCode{0xff,0xff,0xff,0xff}},
		{rv32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, rv32MachineCode{0x00,0x00,0x00,0x00}},
		{rv32{1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, rv32MachineCode{0xff,0x00,0x00,0x00}},
		{rv32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1}, rv32MachineCode{0x00,0x00,0x00,0xff}},
		{rv32{0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}, rv32MachineCode{0x01,0x00,0x00,0x01}},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		output := convert32ByteArrayFrom32bitTo4bit(input)
		correctOutput := testCase[i].output
		if !checkRv32MachineCode(output, correctOutput) {
			t.Errorf("convert32ByteArrayFrom32bitTo4bit(%v) is %v , but need %v. \n", input, output, correctOutput)
		}
	}

}
