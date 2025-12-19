package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func main() {
	const VALID_CHARS = "<>+-.,[]"
	var loop_stack []uint64
	var SM stateMachine = *newStateMachine(100)
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var FS = fs.FS

	if len(os.Args) < 2 {
		panic("No input file provided")
	}

	input_file := os.Args[1]

	input_bytes, err := fs.ReadFile(fs.FS, input_file)
	if err != nil {
		panic(err)
	}

	input := filterString(string(input_bytes), VALID_CHARS)

	for _, instruction := range input {
		switch instruction {
		case '<':
			SM.left()
		case '>':
			SM.right()
		case '+':
			SM.increment()
		case '-':
			SM.decrement()
		case '.':
			SM.output()
		case ',':
			byte, err := reader.ReadByte()
			if err == io.EOF {
				byte = 0
			} else if err != nil {
				panic(fmt.Sprintf("Error reading stdin: %v\n", err))
			}
			SM.input(byte)
		case '[':
			if SM.tape[SM.index] != 0 {
				loop_stack.append(SM.index)
			}
		}

	}

}
