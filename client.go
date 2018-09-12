package main

import (
	"fmt"
	"net"
)

var message string

func main() {
	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connect success ....")

	defer conn.Close()

	for {
		fmt.Println(">Please input data")

		fmt.Scan(&message)

		_, err := conn.Write([]byte(message))

		if err != nil {
			fmt.Println(err)
			return
		}

		data := make([]byte, 1024)

		n, err := conn.Read(data)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(string(data[:n]))
	}
}
