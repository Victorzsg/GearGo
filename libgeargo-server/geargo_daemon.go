package libgeargo_server

import (
	"net"
	"strconv"
	"log"
	"bufio"
	"io"
	"fmt"
)

type Geargo_daemon struct {
	Port int
}

func handleConnection(conn net.Conn) {
	br := bufio.NewReader(conn)
	for{
		data, err := br.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Printf("%s", data)
		fmt.Fprintf(conn, "OK\n")
	}
	conn.Close()
}

func (this *Geargo_daemon) Start() (err error) {

	ln, err := net.Listen("tcp", ":" + strconv.Itoa(this.Port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("get client connection error: ", err)
		}
		go handleConnection(conn)
	}
	return nil
}
