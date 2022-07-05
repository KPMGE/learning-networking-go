package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:3344")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := listener.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Listening...")
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			buffer := make([]byte, 1024)

			for {
				n, err := c.Read(buffer)

				if err != nil {
					if err != io.EOF {
						log.Fatal(err)
					}
					return
				}

				fmt.Printf("read 1024 bytes: %v\n", buffer[:n])
				fmt.Printf("received 1024 bytes: %v\n", string(buffer[:n]))
			}
		}(conn)
	}
}
