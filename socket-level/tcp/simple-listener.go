// extremely simple tcp listener

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// creates a listener with tcp protocol listening on
	// ip address 127.0.0.1 and a random port
	// NOTE: when empty or 0, go will take a random one, the same applies for the ip address
	listener, err := net.Listen("tcp", "127.0.0.1:0")

	// if handkshake process fails, panics
	if err != nil {
		log.Fatal(err)
	}

	// gracefully terminating tcp connection
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// shows the address from which the listener was bounded to
	fmt.Printf("bound to: %v\n", listener.Addr())
}
