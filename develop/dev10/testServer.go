package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// выходить по сигналу ^C
func main() {
	_, cancel := context.WithCancel(context.Background())

	fmt.Println("Starting server ...")
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// handle signal
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		signals := <-sigs
		cancel()
		fmt.Println("\nServer stopped by signal: ", signals)
		os.Exit(0)
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			conn.Close()
			log.Fatal(err)
		}

		fmt.Println("Connection", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	input := make([]byte, 1024)

	for {
		n, err := conn.Read(input)
		if err == io.EOF {
			break
		} else if n == 0 || err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write(append([]byte("from you: "), input[:n]...))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("from connection: ", string(input[:n]))
	}
	fmt.Println("Connection closed")
}
