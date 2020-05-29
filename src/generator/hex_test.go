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
		{[]byte{0x00, 0x00},
			[]byte{58, 48, 50, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 50, 10,
				58, 48, 48, 48, 48, 48, 48, 48, 49, 'F', 'F', 10},
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
