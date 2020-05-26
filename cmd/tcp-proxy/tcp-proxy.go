package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 512)
	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}

		log.Printf("Read %d bytes: %s\n", size, string(b))
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	fmt.Println("hi")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Unable to bind port")
	}
	log.Println("Listening on 0.0.0.0:8080")
	for {
		conn, err := listener.Accept()
		log.Println("Received Connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go echo(conn)

	}

}
