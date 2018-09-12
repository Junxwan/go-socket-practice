package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Start listen")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		go run(conn)
	}
}

func run(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connect :", conn.RemoteAddr())

	for {
		data := make([]byte, 1024)

		n, err := conn.Read(data)

		addr := conn.RemoteAddr()

		if n == 0 {
			fmt.Printf("%s has disconnect\n", addr)
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Receive '%s' from '%s'\n", string(data[:n]), addr)

		_, err = conn.Write([]byte(">success"))

		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
