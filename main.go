package main

import (
	"bufio"
	"fmt"
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

	fmt.Println("Server listening on port 4000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		fmt.Println("Client disconnected:", conn.RemoteAddr())
		conn.Close()
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "/time":
			conn.Write([]byte(time.Now().Format(time.RFC1123) + "\n"))
		case "/quit":
			conn.Write([]byte("Goodbye!\n"))
			return
		default:
			conn.Write([]byte(input + "\n"))
		}
	}
}
