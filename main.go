package main

import (
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/kvault/gbgo/cpu"
)

// GB TODO Document
type GB struct {
	cpu *cpu.CPU
}

/*
func main() {
	gb := GB{
		cpu: cpu.NewCPU(),
	}

	gb.cpu.Run()
}*/

const SockAddr = "/tmp/app.gbgo"

func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())

	const frameSize = 160 * 144 * 10 * 10
	var frame [frameSize]byte
	for i := 0; i < frameSize; i++ {
		frame[i] = byte(i % 8)
	}

	for {
		sendRandomMemoryData(&c)
		time.Sleep(1 * time.Second)
	}
}

func sendRandomMemoryData(c *net.Conn) {
	var msgStream [3]byte

	msgStream[0] = MessageTypes.memoryChanged
	msgStream[1] = byte(rand.Intn(20))
	msgStream[2] = byte(rand.Intn(256))

	(*c).Write(msgStream[0:3])
}

func main() {

	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Printf("Listening")
	defer l.Close()

	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(conn)
	}
}
