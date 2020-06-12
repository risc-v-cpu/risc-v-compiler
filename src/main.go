package main

import (
	"fmt"
	"risc-v-compiler/compiler"
	"risc-v-compiler/generator"
	"risc-v-compiler/util"
)

func main() {
	fmt.Println("RISC-V compiler Version 1.0.0")
	fmt.Println("Author: cw1997 <chagnwei1006@gmail.com>")
	fmt.Println("Repo: https://github.com/risc-v-cpu/risc-v-compiler")
	fmt.Println()
	run("./input.asm", "./output.hex")
}

func run(input string, output string) {
	strArray := util.ReadFileToLinesWithoutLineBreak(input)
	byteArray := compiler.CompileAll(strArray)
	hexArray := generator.GenerateHex(byteArray)
	util.WriteFilesByByteArray(output, hexArray)
}
