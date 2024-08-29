package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer dial.Close()
	mustCopy(os.Stdout, dial)

}

func mustCopy(stdout io.Writer, dial io.Reader) {
	_, err := io.Copy(stdout, dial)
	if err != nil {
		log.Fatal(err)
	}

}
