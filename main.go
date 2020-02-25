package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/kvault/gbgo/pkg/cpu"
	"github.com/kvault/gbgo/pkg/ipc"
)

// GB TODO Document
type GB struct {
	cpu *cpu.CPU
}

func main() {
	_ = GB{
		cpu: cpu.NewCPU(),
	}
	ipc := ipc.New("/tmp/app.gbgo")
	ipc.Start()

	ipc.LogChan <- "Hello World"

	for {
		ipc.LogChan <- "Foo Bar"
		time.Sleep(1 * time.Second)
	}

	for {
		fmt.Println(<-ipc.MemoryChan)
	}
}

func sendRandomMemoryData(c *net.Conn) {
	var msgStream [3]byte

	msgStream[0] = 0
	msgStream[1] = byte(rand.Intn(20))
	msgStream[2] = byte(rand.Intn(256))

	(*c).Write(msgStream[0:3])
}
