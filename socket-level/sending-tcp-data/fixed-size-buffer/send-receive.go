package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	const payload = "Hello client, did u receive my message?"

	listener, err := net.Listen("tcp", "127.0.0.1:5003")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server up and sending payload...")
	go func() {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()

		_, err = conn.Write([]byte(payload))

		if err != nil {
			log.Fatal(err)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:5003")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	buffer := make([]byte, 1024)

	fmt.Println("client waiting for server messages...")
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		fmt.Printf("%v bytes read from server: %v", n, string(buffer[:n]))
	}
}
