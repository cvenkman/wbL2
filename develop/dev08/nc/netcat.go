package nc

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func Netcat(host, port string) error {
	connection, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return err
	}

	r := bufio.NewReader(os.Stdin)
	for {
		// save data form STDIN to data
		data, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// close connection and return if pressed Ctrl+D
				closeConnection(connection)
				return nil
			}
			fmt.Println(err)
		}

		// send (writes) data
		_, err = fmt.Fprintf(connection, "%v", data)
		if err != nil {
			// if server closed connection
			closeConnection(connection)
			return err
		}
	}
}

func closeConnection(connection net.Conn) {
	err := connection.Close()
	if err != nil {
		fmt.Println(err)
	}
}
