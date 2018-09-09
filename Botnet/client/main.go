package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
)

func connection() net.Conn {

	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
	}
	return c

}

func messageHandler(msg string) {

	con := connection()
	err := gob.NewEncoder(con).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

}

func commandProcessor() {

	for {
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if message == "\n" {
			break
		}
		message = strings.TrimSuffix(message, "\n")
		fmt.Println("Sending: ", message)
		messageHandler(message)
	}

}

func main() {

	fmt.Println("Type your message to be sent to the server: ")
	for {
		commandProcessor()
	}

}
