package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"syscall"
)

func passOrDie(err error) {
	if err != nil {
		log.Fatal("Straight to jail")
		log.Fatal(err)
	}
}

func handler(connection net.Conn) {
	var msgLength int32
	err := binary.Read(connection, binary.BigEndian, &msgLength)
	passOrDie(err)

	buf := make([]byte, msgLength)
	err = binary.Read(connection, binary.BigEndian, buf)
	fmt.Println(string(buf))
}

func main() {
	syscall.Unlink("/tmp/brexit.sock")
	listener, err := net.Listen("unix", "/tmp/brexit.sock")
	passOrDie(err)

	for {
		connection, err := listener.Accept()
		passOrDie(err)

		go handler(connection)
	}
}
