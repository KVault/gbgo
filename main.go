package main

import (
	"github.com/kvault/gbgo/pkg/log"
	"math/rand"
	"net"
	"time"

	"github.com/kvault/gbgo/pkg/hardware"
)

// GB TODO Document
type GB struct {
	cpu *hardware.CPU
	rom *hardware.ROM
}

func (gb *GB) insertROM() {
	// from 0x000 to 0x7FFF the cartrige.
	var memoryPointer uint16 = 0x0000

	for i := 0; i < 0x7FFF && i < len(gb.rom.Dump); i++ {
		hardware.Write(memoryPointer, gb.rom.Dump[i])
		memoryPointer++
	}
}

func main() {
	gb := GB{
		cpu: hardware.NewCPU(),
	}

	gb.rom, _ = hardware.LoadROM("./roms/cpu_instrs.gb")
	gb.insertROM()

	for {
		log.Debug("Hi there")
		time.Sleep(500 * time.Millisecond)
	}
}

func sendRandomMemoryData(c *net.Conn) {
	var msgStream [3]byte

	msgStream[0] = 0
	msgStream[1] = byte(rand.Intn(20))
	msgStream[2] = byte(rand.Intn(256))

	(*c).Write(msgStream[0:3])
}
