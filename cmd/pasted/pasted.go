package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"gitlab.0pointer.org/choopm/pasted/pkg/common"
)

var (
	host    = os.Getenv("HOST")
	urlRoot = os.Getenv("URL_ROOT")
	port    = os.Getenv("PORT")
)

func main() {
	rand.Seed(time.Now().UnixNano())

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
	hostAddr, filePath, fileName := common.MakeFileName("/data", conn.RemoteAddr().String())
	os.MkdirAll(filePath, os.ModePerm)
	f, err := os.Create(filePath + fileName)
	check(err)

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
		check(err)

		// Read again
		reqLen, err = conn.Read(buf)
	}
	conn.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
