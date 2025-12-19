package main

type gopherfuck struct {
	tape  []uint8
	index uint
}

func (self *gopherfuck) increment() {
	self.tape[self.index]++
}

func (self *gopherfuck) decrement() {
	self.tape[self.index]--
}

func (self *gopherfuck) right() {
	self.index++
}

func (self *gopherfuck) left() {
	self.index--
}
