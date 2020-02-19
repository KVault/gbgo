package cpu

import "fmt"

// CPU Document this
type CPU struct {
}

// Run Infinite loop containing the basic CPU workflow. Which is decode - execute and repeat until death
func (cpu CPU) Run() {
	for {
		fmt.Println(RAM.bank)
	}
}

// NewCPU Creates a new CPU with some default values, such as Speed, memory and Opcodes
func NewCPU() *CPU {
	c := &CPU{}
	return c
}
