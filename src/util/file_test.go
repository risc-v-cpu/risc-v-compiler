package util

import (
	"strings"
	"testing"
)

type ReadFileToLinesWithoutLineBreakTestCase struct {
	path   string
	output []string
}

func TestReadFileToLines(t *testing.T) {
	testCase := []ReadFileToLinesWithoutLineBreakTestCase{
		{"./file_test_case/risc-v.txt", []string{"addi a1,a2,a3", "sub a1,a2,a3"}},
		{"./file_test_case/123456789.txt", []string{"123", "456", "789"}},
	}

	for i := 0; i < len(testCase); i++ {
		path := testCase[i].path
		output := ReadFileToLinesWithoutLineBreak(path)
		correctOutput := testCase[i].output
		if strings.Join(output, "\n") != strings.Join(correctOutput, "\n") {
			t.Errorf("ReadFileToLinesWithoutLineBreak(%v) is %v , but need %v. \n", path, output, correctOutput)
		}
	}

}

func compareStringArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for k, _ := range a {
		if a[k] != b[k] {
			return false
		}
	}

	return true
}

type WriteFilesByByteArrayTestCase struct {
	input   WriteFilesByByteArrayInput
	output []string
}

type WriteFilesByByteArrayInput struct {
	path string
	byteArray []byte
}

func TestWriteFilesByByteArray(t *testing.T) {
	testCase := []WriteFilesByByteArrayTestCase{
		{WriteFilesByByteArrayInput{"./file_test_case/write_test.txt", []byte{65,66,67,68,69}}, []string{"ABCDE"}},
		{WriteFilesByByteArrayInput{"./file_test_case/write_test.txt", []byte{'0','1','2','3','4'}}, []string{"01234"}},
	}

	for i := 0; i < len(testCase); i++ {
		input := testCase[i].input
		path := input.path
		byteArray := input.byteArray
		WriteFilesByByteArray(path, byteArray)
		output := ReadFileToLinesWithoutLineBreak(path)
		correctOutput := testCase[i].output
		if !compareStringArray(output, correctOutput) {
			t.Errorf("ReadFileToLinesWithoutLineBreak(%v) is %v , but need %v. \n", path, output, correctOutput)
		}
	}

}
