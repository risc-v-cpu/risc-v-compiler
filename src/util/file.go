package util

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadFileToLinesWithoutLineBreak(path string) []string {
	byteArray, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Read file error, %v Please check the path %s . \n", err, path)
	}

	str := string(byteArray)
	str = strings.ReplaceAll(str, "\r\n", "\n")
	str = strings.ReplaceAll(str, "\r", "\n")

	rawLines := strings.Split(str, "\n")
	length := len(rawLines)

	var result []string

	for i := 0; i < length; i++ {
		if rawLines[i] != "" {
			result = append(result, rawLines[i])
		}
	}

	return result
}

func WriteFilesByByteArray(path string, byteArray[]byte) {
	err := ioutil.WriteFile(path, byteArray, os.FileMode(777))
	if err != nil {
		log.Fatalf("Write file error, %v Please check the path %s and content. \n", err, path)
	}
}
