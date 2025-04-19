package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server listening on :4000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Printf("Client connected: %s at %s\n", addr, time.Now().Format(time.RFC3339))
	defer fmt.Printf("Client disconnected: %s at %s\n", addr, time.Now().Format(time.RFC3339))

	buffer := make([]byte, 0, 1024)
	temp := make([]byte, 1)

	for {
		n, err := conn.Read(temp)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client closed the connection.")
			} else {
				fmt.Println("Error reading from client:", err)
			}
			return
		}
		if n > 0 {
			if temp[0] == '\n' {
				message := strings.TrimSpace(string(buffer))
				fmt.Printf("Received: %s\n", message)
				conn.Write([]byte(message + "\n"))
				buffer = buffer[:0] // clear buffer
			} else {
				buffer = append(buffer, temp[0])
			}
		}
	}
}