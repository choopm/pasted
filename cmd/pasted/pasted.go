package main

import (
	"fmt"
	"io"
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

	bytes, err := io.Copy(f, conn)
	check(err)

	fmt.Println(bytes, "bytes from", conn.RemoteAddr().String()+":", url)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
