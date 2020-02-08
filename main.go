package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:4444")
	if err != nil {
		fmt.Errorf("Error when connecting to server: %v\n", err)
	}
	//for {
	input := []byte("Hello World!")
	_, err = conn.Write(input)
	if err != nil {
		fmt.Errorf("Error when sending message: %v\n", err)
	}
	//}
}
