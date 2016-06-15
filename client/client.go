package main

import (
	"encoding/binary"
	"net"
)

func main() {
	socket, _ := net.Dial("unix", "/tmp/brexit.sock")
	defer socket.Close()
	str := []byte("vote remain")
	msgLength := int32(len(str))
	binary.Write(socket, binary.BigEndian, msgLength)
	socket.Write(str)
}
