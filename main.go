package main

import (
	"fmt"
	"net"
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
		go handleConnection(conn) //concurrent
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Printf("Client connected: %s at %s\n", addr, time.Now().Format(time.RFC3339))
	defer fmt.Printf("Client disconnected: %s at %s\n", addr, time.Now().Format(time.RFC3339))

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Error writing to client:", err)
		}
	}
}