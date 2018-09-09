package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(netadd string) {
	ln, err := net.Listen("tcp", netadd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening: ", ln.Addr())

	for {
		c, err := ln.Accept()
		fmt.Println("New message from client! Address: ", ln.Addr())
		if err != nil {
			fmt.Println(err)
			continue
		}
		connectionHandler(c)
		defer c.Close()
	}

}

func connectionHandler(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Recieved: ", msg)
	}
	c.Close()
}

func main() {
	fmt.Println("Enter address and port to listen on. (Example: 127.0.0.1:9999)")
	var netadd string
	fmt.Scanln(&netadd)
	go server(netadd)
	var input string
	fmt.Scanln(&input)
}
