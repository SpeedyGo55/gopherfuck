package main

import "fmt"

type stateMachine struct {
	tape  []byte
	index uint64
}

func (self *stateMachine) increment() {
	self.tape[self.index]++
}

func (self *stateMachine) decrement() {
	self.tape[self.index]--
}

func (self *stateMachine) right() {
	self.index++
	if self.index >= uint64(len(self.tape)) {
		self.index = 0
	}
}

func (self *stateMachine) left() {
	if self.index == 0 {
		self.index = uint64(len(self.tape) - 1)
	} else {
		self.index--
	}

}

func (self *stateMachine) output() {
	fmt.Printf("%c", self.tape[self.index])
}

func (self *stateMachine) input(in byte) {
	self.tape[self.index] = in
}

func newStateMachine(length int64) *stateMachine {
	sm := stateMachine{tape: make([]byte, length), index: 0}
	return &sm
}
