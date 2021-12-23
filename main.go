package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9001"
	}

	addr := fmt.Sprintf("0.0.0.0:%s", PORT)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server is running at %s ...", addr)
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}

		welcome := fmt.Sprintf("Hello %s!\n", conn.RemoteAddr().String())
		conn.Write([]byte(welcome))
	}
}
