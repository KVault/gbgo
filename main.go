package main

import (
	"math/rand"
	"net"

	"github.com/kvault/gbgo/pkg/hardware"
	"github.com/kvault/gbgo/pkg/ipc"
)

// GB TODO Document
type GB struct {
	cpu *hardware.CPU
	rom *hardware.ROM
}

func (gb *GB) insertROM(){
	// from 0x000 to 0x7FFF the cartrige.
	var memoryPointer uint16 = 0x0000

	for i := 0; i < 0x7FFF && i < len(gb.rom.Dump); i++{
		hardware.Write(memoryPointer, gb.rom.Dump[i])
		memoryPointer++
	}
}

func main() {
	gb := GB{
		cpu: hardware.NewCPU(),
	}

	gb.rom, _ = hardware.LoadROM("./roms/cpu_instrs.gb");
	gb.insertROM()

	ipc := ipc.New("/tmp/app.gbgo")
	ipc.Start()

	for i, content := range hardware.RAM.Bank {
		var msgStream [2]byte

		msgStream[0] = byte(i)
		msgStream[1] = content

		ipc.MemoryChan <- msgStream[0:]
	}

	for {

	}
}

func sendRandomMemoryData(c *net.Conn) {
	var msgStream [3]byte

	msgStream[0] = 0
	msgStream[1] = byte(rand.Intn(20))
	msgStream[2] = byte(rand.Intn(256))

	(*c).Write(msgStream[0:3])
}
