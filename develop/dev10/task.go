package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
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

func main() {
	var timeout int
	flag.IntVar(&timeout, "timeout", 10, "timeout")

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet --timeout=10s host port")
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	// connects to the address on the named network (tcp)
	// each IP address given the time to connect
	// JoinHostPort combines host and port into a network address of the form "host:port"
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Duration(timeout)*time.Second)
	if err != nil {
		// address doesn't exist
		time.Sleep(time.Duration(timeout) * time.Second)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":8080", nil)
}

// http://localhost:8080
