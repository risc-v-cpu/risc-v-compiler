package compiler

import (
	"testing"
)

type setBitForRv32TestCase struct {
	input  setBitForRv32Input
	output rv32
}

type setBitForRv32Input struct {
	raw rv32
	strCode string
	startIndex int
}

func TestSetBitForRv32(t *testing.T) {
	testCase := []setBitForRv32TestCase{
		{ setBitForRv32Input{
			rv32{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, "1111111", 0},
			rv32{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		},
		{ setBitForRv32Input{
			rv32{1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, "101", 3},
			rv32{1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0},
		},
		{ setBitForRv32Input{
			rv32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1}, "00000000", 0},
			rv32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		},
		{ setBitForRv32Input{
			rv32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}, "1111111", 0},
			rv32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1},
		},
		{ setBitForRv32Input{
			rv32{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, "01111111", 1},
			rv32{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,0},
		},
		{ setBitForRv32Input{
			rv32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, "110100100110", 20},
			rv32{1,1,0,1,0,0,1,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		input.raw.setBitForRv32(input.strCode, input.startIndex)
		output := input.raw
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("raw.setBitForRv32(%v) is %v , but need %v. \n", input, output, correctOutput)
		}
	}

}


type setOpCodeTestCase struct {
	input setOpCodeInput
	output rv32
}

type setOpCodeInput struct {
	raw rv32
	strCode string
}

func TestSetOpCode(t *testing.T) {
	testCase := []setOpCodeTestCase{
		{setOpCodeInput{
			rv32{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, "1110010"},
			rv32{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,0,0,1,0},
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		input.raw.setOpCode(input.strCode)
		output := input.raw
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("raw.setOpCode(%v) is %v , but need %v. \n", input, output, correctOutput)
		}
	}

}


type setRdTestCase struct {
	input setRdInput
	output rv32
}

type setRdInput struct {
	raw rv32
	rd int
}

func TestSetRd(t *testing.T) {
	testCase := []setRdTestCase{
		{setRdInput{
			rv32{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,0,0,1,0}, 2},
			rv32{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0},
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		input.raw.setRd(input.rd)
		output := input.raw
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("raw.setRd(%v) is %v , but need %v. \n", input, output, correctOutput)
		}
	}

}


type setFunct3TestCase struct {
	input setFunct3Input
	output rv32
}

type setFunct3Input struct {
	raw rv32
	strCode string
}

func TestSetFunct3(t *testing.T) {
	testCase := []setFunct3TestCase{
		{setFunct3Input{
			rv32{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0}, "101"},
			rv32{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,1,0,1,1,1,0,0,1,0},
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		input.raw.setFunct3(input.strCode)
		output := input.raw
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("raw.setFunct3(%v) is %v , but need %v. \n", input, output, correctOutput)
		}
	}

}


type setRs1TestCase struct {
	input setRs1Input
	output rv32
}

type setRs1Input struct {
	raw rv32
	rs1 int
}

func TestSetRs1(t *testing.T) {
	testCase := []setRs1TestCase{
		{setRs1Input{rv32{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,1,0,1,1,1,0,0,1,0}, 4},
			rv32{1,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,1,0,0,0,1,0,1,1,1,0,0,1,0},
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		input.raw.setRs1(input.rs1)
		output := input.raw
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("raw.setRs1(%v) is %v , but need %v. \n", input, output, correctOutput)
		}
	}

}


type setRs2TestCase struct {
	input setRs2Input
	output rv32
}

type setRs2Input struct {
	raw rv32
	rs2 int
}

func TestSetRs2(t *testing.T) {
	testCase := []setRs2TestCase{
		{setRs2Input{
			rv32{1,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,1,0,0,0,1,0,1,1,1,0,0,1,0}, 8},
			rv32{1,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,0,1,0,1,1,1,0,0,1,0},
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		input.raw.setRs2(input.rs2)
		output := input.raw
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("raw.setRs2(%v) is %v , but need %v. \n", input, output, correctOutput)
		}
	}

}


type setFunct7TestCase struct {
	input setFunct7Input
	output rv32
}

type setFunct7Input struct {
	raw rv32
	strCode string
}

func TestSetFunct7(t *testing.T) {
	testCase := []setFunct7TestCase{
		{setFunct7Input{
			rv32{1,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,0,1,0,1,1,1,0,0,1,0}, "0100000"},
			rv32{0,1,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,1,0,0,0,1,0,1,1,1,0,0,1,0},
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		input.raw.setFunct7(input.strCode)
		output := input.raw
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("raw.setFunct7(%v) is %v , but need %v. \n", input, output, correctOutput)
		}
	}

}


type setImmBitForRv32TypeITestCase struct {
	input setImmBitForRv32TypeIInput
	output rv32
}

type setImmBitForRv32TypeIInput struct {
	raw rv32
	imm int
}

func TestSetImmBitForRv32TypeI(t *testing.T) {
	testCase := []setImmBitForRv32TypeITestCase{
		// 1101 0011 1011‬ = 0xd3b = 3387
		{setImmBitForRv32TypeIInput{
			  rv32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,1,0,0,0,1,0,1,1,1,0,0,1,0}, 0xd3b},
			rv32{1,1,0,1,0,0,1,1,1,0,1,1,0,0,1,0,0,1,0,1,0,0,0,1,0,1,1,1,0,0,1,0},
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		input.raw.setImmBitForRv32TypeI(input.imm)
		output := input.raw
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("raw.setImmBitForRv32TypeI(%v) is %v , but need %v. \n", input, output, correctOutput)
		}
	}

}


type setImmBitForRv32TypeSTestCase struct {
	input setImmBitForRv32TypeSInput
	output rv32
}

type setImmBitForRv32TypeSInput struct {
	raw rv32
	imm int
}

func TestSetImmBitForRv32TypeS(t *testing.T) {
	testCase := []setImmBitForRv32TypeSTestCase{
		// 1101 1011 0010‬ = 0xdb2 = 3506
		{setImmBitForRv32TypeSInput{
			  rv32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,1,0,0,0,0,0,1,1,1,0,0,1,0}, 0xdb2},
			rv32{1,1,0,1,1,0,1,0,0,0,0,0,0,0,1,0,0,1,0,1,1,0,0,1,0,1,1,1,0,0,1,0},
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		input.raw.setImmBitForRv32TypeS(input.imm)
		output := input.raw
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("raw.setImmBitForRv32TypeS(%v) is \n%v , but need \n%v. \n", input, output, correctOutput)
		}
	}

}


type setImmBitForRv32TypeBTestCase struct {
	input setImmBitForRv32TypeBInput
	output rv32
}

type setImmBitForRv32TypeBInput struct {
	raw rv32
	imm int
}

func TestSetImmBitForRv32TypeB(t *testing.T) {
	testCase := []setImmBitForRv32TypeBTestCase{
		// 1 1101 1011 0010‬ = 0x1db2 = 7602
		{setImmBitForRv32TypeBInput{
			rv32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,1,0,0,0,0,0,1,1,1,0,0,1,0}, 0x1db2},
			rv32{1,1,0,1,1,0,1,0,0,0,0,0,0,0,1,0,0,1,0,1,1,0,0,1,1,1,1,1,0,0,1,0},
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		input.raw.setImmBitForRv32TypeB(input.imm)
		output := input.raw
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("raw.setImmBitForRv32TypeB(%v) is \n%v , but need \n%v. \n", input, output, correctOutput)
		}
	}

}


type setImmBitForRv32TypeUTestCase struct {
	input setImmBitForRv32TypeUInput
	output rv32
}

type setImmBitForRv32TypeUInput struct {
	raw rv32
	imm int
}

func TestSetImmBitForRv32TypeU(t *testing.T) {
	testCase := []setImmBitForRv32TypeUTestCase{
		// 1101 1011 0010 1101 1011_ 0010 1101 1011‬ = 0xdb2 = 3,677,205,211‬
		{setImmBitForRv32TypeUInput{
			rv32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,0,0,1,0}, 0xdb2db2db},
			rv32{1,1,0,1,1,0,1,1,0,0,1,0,1,1,0,1,1,0,1,1,0,0,0,0,0,1,1,1,0,0,1,0},
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		input.raw.setImmBitForRv32TypeU(input.imm)
		output := input.raw
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("raw.setImmBitForRv32TypeU(%v) is \n%v , but need \n%v. \n", input, output, correctOutput)
		}
	}

}
