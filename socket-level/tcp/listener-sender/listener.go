package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:3344")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Listening on 127.0.0.1:3344...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer conn.Close()

			scanner := bufio.NewScanner(conn)
			scanner.Split(bufio.ScanWords)

			var words []string

			for scanner.Scan() {
				words = append(words, scanner.Text())
			}

			err = scanner.Err()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("message: %v\n", words)
		}(conn)
	}
}
