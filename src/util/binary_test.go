package util

import (
	"testing"
)

type IntToBinaryStringTestCase struct {
	num      int
	zeroFill int
	output   string
}

func TestConvertToBin(t *testing.T) {
	testCase := []IntToBinaryStringTestCase{
		{0, 8, "00000000"},
		{1, 8, "00000001"},
		{8, 8, "00001000"},
		{6, 6, "000110"},
		{4, 1, "100"},
	}

	for i := 0; i < len(testCase); i++ {
		num := testCase[i].num
		zeroFill := testCase[i].zeroFill
		output := IntToBinaryString(num, zeroFill)
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("IntToBinaryString(%v, %v) is %v , but need %v. \n", num, zeroFill, output, correctOutput)
		}
	}

}
