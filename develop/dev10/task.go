package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port: go-telnet mysite.ru 8080; go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

//https://www.opennet.ru/man.shtml?topic=telnet&category=1&russian=0
//https://pkg.go.dev/net

const usage = "Usage: go-telnet --timeout=10s host port"

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "timeout")
	flag.Parse()

	// get host and port
	if flag.NArg() < 2 {
		log.Fatal(usage)
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	fmt.Printf("Trying %s...\n", host)
	connection := createConnection(host, port, timeout)
	// close tcp connection
	defer closeConnection(connection)
	fmt.Println("Connected")

	r := bufio.NewReader(os.Stdin)
	for {
		// save data form STDIN to data
		data, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// close connection and return if pressed Ctrl+D
				closeConnection(connection)
				return
			}
			fmt.Println(err)
		}

		// send (writes) data
		_, err = fmt.Fprintf(connection, "%v", data)
		if err != nil {
			// if server closed connection
			fmt.Println(err)
			closeConnection(connection)
			return
		}
	}
}

// connects to the address (host:port) on the tcp
func createConnection(host string, port string, timeout time.Duration) net.Conn {
	// each IP address given the time to connect
	// JoinHostPort combines host and port into a network address of the form "host:port"
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		// address doesn't exist
		// exit after timeout
		time.Sleep(time.Duration(timeout) * time.Second)
		fmt.Println(err)
		os.Exit(1)
	}
	return connection
}

func closeConnection(connection net.Conn) {
	err := connection.Close()
	if err != nil {
		log.Fatal(err)
	}
}
