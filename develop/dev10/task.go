package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	// "net/http"
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

	if flag.NArg() < 2 {
		log.Fatal(usage)
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	fmt.Println(host, port, net.JoinHostPort(host, port))

	// connects to the address on the named network (tcp)
	// each IP address given the time to connect
	// JoinHostPort combines host and port into a network address of the form "host:port"
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		// address doesn't exist
		time.Sleep(time.Duration(timeout) * time.Second)
		os.Exit(1)
	}

	// close tcp connection
	defer func() {
		err := connection.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go readOutConn(connection)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		_, err := fmt.Fprintf(connection, s.Text()+" / HTTP/1.0\r\n\r\n")
		if err != nil {
			fmt.Println(err)
			break
		}
	}

}

func readOutConn(connection net.Conn) {
	r := bufio.NewReader(connection)
	for {
		message, err := r.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Connection closed by foreign host")
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print(message)
	}
}

// http://localhost:8080
