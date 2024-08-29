package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		return
	}

	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(accept)
	}

}

func handleConn_01(accept net.Conn) {
	defer accept.Close()
	for {
		_, err := io.WriteString(accept, time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
