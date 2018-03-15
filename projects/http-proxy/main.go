package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	host := flag.String("host", "", "host")
	port := flag.String("port", "8080", "port")
	flag.Parse()

	listener, err := net.Listen("tcp", *host+":"+*port)

	if err != nil {
		log.Fatal(err)
	}

	conns := make(chan net.Conn)

	go handleConns(listener, conns)
	for {
		go handleRequest(<-conns)
	}
}

func handleConns(listener net.Listener, conns chan net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		conns <- conn
	}
}

func handleRequest(conn net.Conn) {

}
