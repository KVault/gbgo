package main

import "github.com/kvault/gbgo/cpu"

// GB TODO Document
type GB struct {
	cpu *cpu.CPU
}

func main() {
	gb := GB{
		cpu: cpu.NewCPU(),
	}

	gb.cpu.Run()
}
