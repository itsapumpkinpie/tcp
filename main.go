package main 

import (
	"io"
	"log"
	"net"
)

func main() {
	// this tells us we can accept and get back a connection and an error when we're listening
	lis, err := net.Listen("tcp", ":8080")
	
	if err != nil {
		log.Panic(err)
	}

	// we also can close our listener
	defer lis.Close()

	// conn is the the reader and the writer.
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHi form the local TCP server!! :)\n")
		conn.Close()
	}
}