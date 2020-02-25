package ipc

import (
	"log"
	"net"
	"os"

	"github.com/kvault/gbgo/pkg"
)

// IPC Represents all the channels that can be communicated between the emulator and the UI
type IPC struct {
	address    string
	connection *net.Conn
	LogChan    chan string
	MemoryChan chan []byte
}

// Start Opens a new connection to the IPC's address.
func (ipc *IPC) Start() error {
	if err := os.RemoveAll(ipc.address); err != nil {
		log.Fatal(err)
		return err
	}

	l, err := net.Listen("unix", ipc.address)
	if err != nil {
		log.Fatal("listen error:", err)
		return err
	}

	// Rinse and repeat
	go ipc.acceptConnections(l)
	go ipc.watchChan()

	return nil
}

func (ipc *IPC) acceptConnections(l net.Listener) error {

	defer l.Close()

	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		ipc.connection = &conn

		go ipc.reader()
	}
}

// Listens to messages over the socker, redirects them to the local channels to be picked up by
// the channel listener. They'll know what to do with it!
func (ipc *IPC) reader() {
	buf := make([]byte, 160*144*2) // Size of a GB frame times 2
	for {
		n, err := (*ipc.connection).Read(buf[:])
		if err != nil {
			return
		}

		// buf[0] defines the type of the message
		// buf[1:n] is the payload
		switch buf[0] {
		case pkg.MemoryUpdated:
			ipc.MemoryChan <- buf[1:n]
			break
		case 48:
			ipc.MemoryChan <- buf[0:2]
			break
		}
	}
}

func (ipc *IPC) watchChan() {
	for {
		select {
		case log := <-ipc.LogChan:
			ipc.sendMessage(pkg.Log, []byte(log))
		
		case memPos := <- ipc.MemoryChan:
			ipc.sendMessage(pkg.MemoryUpdated, memPos)
		}
	}
}

func (ipc *IPC) sendMessage(action int, payload []byte) {
	if(ipc.connection != nil){	
		actionPrefix := []byte{byte(action)}
		message := append(actionPrefix, payload...)
		(*ipc.connection).Write(message[0:])
	}
}

func echoServer(c net.Conn) {
	const frameSize = 160 * 144 * 10 * 10
	var frame [frameSize]byte
	for i := 0; i < frameSize; i++ {
		frame[i] = byte(i % 8)
	}
	/*
		for {
			sendRandomMemoryData(&c)
			time.Sleep(1 * time.Second)
		}
	*/
}

// Stop cleans up and closes all the open connections.
func (ipc *IPC) Stop() error {
	return nil
}

// New Builds a new IPC instance
func New(address string) *IPC {
	return &IPC{
		address:    address,
		MemoryChan: make(chan []byte),
		LogChan:    make(chan string),
	}
}
