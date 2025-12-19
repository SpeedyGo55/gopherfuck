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
}

func (self *stateMachine) left() {
	self.index--
}

func (self *stateMachine) jump(index uint) {
	self.index = index
}

func (self *stateMachine) output() {
	fmt.Printf("Test %c\n", self.tape[self.index])
}

func (self *stateMachine) input(in byte) {
	self.tape[self.index] = in
}

func newStateMachine(length int64) *stateMachine {
	sm := stateMachine{tape: make([]byte, length), index: 0}
	return &sm
}
