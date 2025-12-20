package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	const VALID_CHARS = "<>+-.,[]"
	var loop_stack []int
	var SM stateMachine = *newStateMachine(30000)
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	if len(os.Args) < 2 {
		panic("No input file provided")
	}
	input_file := os.Args[1]

	input_bytes, err := os.ReadFile(input_file)
	if err != nil {
		panic(err)
	}
	input := filterString(string(input_bytes), VALID_CHARS)

	// Validate brackets first
	if !validateBrackets(input) {
		panic("Unmatched brackets in program")
	}

	// Use manual index control instead of range
	for instruction_index := 0; instruction_index < len(input); instruction_index++ {
		instruction := input[instruction_index]

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
			if SM.tape[SM.index] == 0 {
				depth := 1
				for i := instruction_index + 1; i < len(input); i++ {
					if input[i] == '[' {
						depth++
					} else if input[i] == ']' {
						depth--
						if depth == 0 {
							instruction_index = i
							break
						}
					}
				}
			} else {
				loop_stack = append(loop_stack, instruction_index)
			}
		case ']':
			if SM.tape[SM.index] != 0 {
				if len(loop_stack) == 0 {
					panic(fmt.Sprintf("Unmatched ] at position %d", instruction_index))
				}
				instruction_index = loop_stack[len(loop_stack)-1]
			} else {
				if len(loop_stack) > 0 {
					loop_stack = loop_stack[:len(loop_stack)-1]
				}
			}
		}
	}
}
