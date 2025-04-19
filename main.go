package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"
	"os"
	"path/filepath"
	"flag"
)

func main() {
    port := flag.String("port", "4000", "Port to listen on")
    flag.Parse()

    listener, err := net.Listen("tcp", ":"+*port)
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    fmt.Printf("Server listening on :%s\n", *port)
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

    addr := conn.RemoteAddr().(*net.TCPAddr).IP.String()
    logFile := filepath.Join(".", addr+".log")
    file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening log file:", err)
        return
    }
    defer file.Close()

    fmt.Printf("Client connected: %s at %s\n", addr, time.Now().Format(time.RFC3339))
    defer fmt.Printf("Client disconnected: %s at %s\n", addr, time.Now().Format(time.RFC3339))

    buffer := make([]byte, 0, 1024)
    temp := make([]byte, 1)

    for {
        // Set 30-second deadline for Read operations
        conn.SetReadDeadline(time.Now().Add(30 * time.Second))

        n, err := conn.Read(temp)
        if err != nil {
            if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
                fmt.Println("Client inactive (30s timeout):", addr)
                conn.Write([]byte("Connection closed due to inactivity.\n"))
            } else if err == io.EOF {
                fmt.Println("Client closed the connection:", addr)
            } else {
                fmt.Println("Error reading from client:", err)
            }
            return
        }

        if n > 0 {
            if temp[0] == '\n' {
                message := strings.TrimSpace(string(buffer))
                fmt.Printf("Received from %s: %s\n", addr, message)
                file.WriteString(fmt.Sprintf("[%s] %s\n", time.Now().Format(time.RFC3339), message))
                conn.Write([]byte(message + "\n"))
                buffer = buffer[:0]
            } else {
                buffer = append(buffer, temp[0])
            }
        }
    }
}