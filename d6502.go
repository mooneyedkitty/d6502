package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mooneyedkitty/cpu6502"
)

// ----------------------------------------------------------------------------
// d6502.go
// 6502 Disassembler
// ----------------------------------------------------------------------------
// Copyright (c) 2024 Robert L. Snyder <rob@mooneyedkitty.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.
// ----------------------------------------------------------------------------

func main() {

	var fileName string
	var startAddress int
	var fileOffset int

	fmt.Println("6502 Disassembler (d6502)")
	fmt.Println("(c) 2024 Rob Snyder <rob@mooneyedkitty.com>")
	fmt.Println("")

	flag.StringVar(&fileName, "f", "", "File to disassemble")
	flag.IntVar(&startAddress, "s", 0, "Starting address (defaults to 0)")
	flag.IntVar(&fileOffset, "o", 0, "Offset in the file to start at (defaults to 0)")

	flag.Parse()

	fmt.Printf("fileName: %v\n", fileName)

	if fileName == "" {
		fmt.Println("File name must be supplied.")
		flag.Usage()
		os.Exit(1)
	}

	cpu6502.Disassemble(&cpu6502.DisassembleOptions{
		FileName:     fileName,
		StartAddress: startAddress,
		FileOffset:   fileOffset,
	})

}
