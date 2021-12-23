package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func handle(c net.Conn) {
	defer c.Close()

	br := bufio.NewReader(c)

	remoteAddr := c.RemoteAddr().String()

	for {
		fmt.Fprintf(c, "> ")

		line, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				continue
			}
			log.Println(err)
		}

		line = strings.TrimSpace(line)
		log.Printf("%s > %s\n", remoteAddr, line)

		switch line {
			case "ip":
				fmt.Fprintf(c, "%s\n", remoteAddr)
			case "rng":
				fmt.Fprintf(c, "%d\n", rand.Int63())
			case "time":
				fmt.Fprintf(c, "%s\n", time.Now())
			default:
				fmt.Fprintf(c, "commands: ip|rng|time\n")
		}
	}
}

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
			log.Println(err)
		}

		welcome := fmt.Sprintf("Hello %s!\n", conn.RemoteAddr().String())
		conn.Write([]byte(welcome))

		go handle(conn)
	}
}
