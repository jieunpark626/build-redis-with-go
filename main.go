package main

import (
	"fmt"
	"net"
	"io"
	"os"
)

func main(){
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()

	for {
		buffer := make([]byte, 1024)

		_, err = connection.Read(buffer)
		if err != nil{
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client : ", err.Error())
			os.Exit(1)
		}

		connection.Write([]byte("+OK\r\n"))
	}
}