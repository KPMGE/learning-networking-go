package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func getInput() string {
	fmt.Print("Enter your text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:5003")

	if err != nil {
		log.Fatal(err)
	}

	for {
		input := getInput()

		go func() {
			conn, err := listener.Accept()

			if err != nil {
				log.Fatal(err)
			}

			defer conn.Close()

			_, err = conn.Write([]byte(input))

			if err != nil {
				log.Fatal(err)
			}
		}()
	}
}
