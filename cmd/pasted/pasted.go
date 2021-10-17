package main

import (
	"fmt"
	"net"
	"os"

	"github.com/choopm/pasted/pkg/common"
)

var (
	host    = os.Getenv("HOST")
	urlRoot = os.Getenv("URL_ROOT")
	port    = os.Getenv("PORT")
)

func main() {
	// Setup listener
	listener, err := net.Listen("tcp", host+":"+port)
	check(err)
	defer listener.Close()
	fmt.Println("Listening on " + host + ":" + port)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		check(err)
		// Handle connection in a goroutine.
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	filePath, fileName := common.MakeFileName("/data", conn.RemoteAddr().String())
	os.MkdirAll(filePath, os.ModePerm)

	f, err := os.Create(filePath + fileName)
	check(err)
	defer f.Close()

	url := urlRoot + fileName
	conn.Write([]byte("Connection established, your paste will be at: " + url + "\n"))
	fmt.Println("New paste from " + conn.RemoteAddr().String() + ": " + url)

	bufSize := 1024
	buf := make([]byte, bufSize)
	for reqLen, err := conn.Read(buf); reqLen > 0; {
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			break
		}

		// Append to file
		_, err := f.Write(buf[0:reqLen])
		if err != nil {
			fmt.Println("Error writing: ", err.Error())
			break
		}

		// Read again
		reqLen, err = conn.Read(buf)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
