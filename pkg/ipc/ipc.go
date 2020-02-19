package ipc

import (
	"log"
	"net"
	"os"
)

type IPC struct {
	address  string
	listener *net.Listener
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
	go acceptConnections(l)

	return nil
}

func acceptConnections(l net.Listener) error {

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

func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())

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
		address: address,
	}
}
