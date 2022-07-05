package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	message := os.Args[1]

	if message == "" {
		fmt.Printf("usage executable <message>")
	}

	conn, err := net.Dial("tcp", "127.0.0.1:3344")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	conn.Write([]byte(message))
}
