package util

import "testing"

type TestReverseTestCase struct {
	input string
	output string
}

func TestReverseString(t *testing.T) {
	testCase := []TestReverseTestCase{
		{"abcdefg", "gfedcba"},
		{"昌维", "维昌"},
		{"123", "321"},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		output := ReverseString(input)
		correctOutput := testCase[i].output
		if output != correctOutput {
			t.Errorf("ReverseString(%v) is \n%v , but need \n%v. \n", input, output, correctOutput)
		}
	}

}
