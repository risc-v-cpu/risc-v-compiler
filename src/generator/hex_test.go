package generator

import (
	"bytes"
	"testing"
)

type GenerateHexTestCase struct {
	input  []byte
	output []byte
}

func TestGenerateHex(t *testing.T) {
	testCase := []GenerateHexTestCase{
		{[]byte{0,16,96,147,0,0,97,19,1,160,97,147,0,16,144,147,0,16,32,35,0,17,1,19,254,49,24,227,254,49,0,227},
			[]byte(":0400000000106093F9\n:040001000000611387\n:0400020001A0619365\n:0400030000109093C6\n:0400040000102023A5\n:0400050000110113D2\n:04000600FE3118E3CC\n:04000700FE3100E3E3\n:00000001FF"),
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		output := GenerateHex(input)
		correctOutput := testCase[i].output
		if bytes.Compare(output, correctOutput) != 0 {
			t.Errorf("GenerateHex(%v) is %v , but need %v. \n", input, output, correctOutput)
		}
	}

}


type getHexCheckTestCase struct {
	input  [][]byte
	output uint8
}

func TestGetHexCheck(t *testing.T) {
	testCase := []getHexCheckTestCase{
		{[][]byte{{4}, {0, 0}, {0}, {0, 16, 96, 147}},
			249,
		},
		{[][]byte{{4}, {0, 1}, {0}, {0, 0, 97, 19}},
			135,
		},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		output := getHexCheck(input...)
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("GenerateHex(%v) is %v , but need %v. \n", input, output, correctOutput)
		}
	}

}
