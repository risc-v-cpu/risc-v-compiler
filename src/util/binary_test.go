package util

import (
	"testing"
)

type IntToBinaryStringTestCase struct {
	num      int
	fill int
	output   string
}

func TestConvertToBin(t *testing.T) {
	testCase := []IntToBinaryStringTestCase{
		{0, 8, "00000000"},
		{1, 8, "00000001"},
		{8, 8, "00001000"},
		{6, 6, "000110"},
		{4, 12, "000000000100"},
		{2, 12, "000000000010"},
		{-2, 12, "111111111110"},
		{-3, 12, "111111111101"},
		{-4, 12, "111111111100"},
		{-8, 12, "111111111000"},
	}

	for i := 0; i < len(testCase); i++ {
		num := testCase[i].num
		fill := testCase[i].fill
		output := IntToBinaryString(num, fill)
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("IntToBinaryString(%v, %v) is %v , but need %v. \n", num, fill, output, correctOutput)
		}
	}

}
