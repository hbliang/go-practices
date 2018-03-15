package main

import (
	"flag"
	"io"
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
	defer listener.Close()

	log.Println("listen on " + *host + ":" + *port)

	conns := make(chan net.Conn)

	go listen(listener, conns)

	for {
		go handleRequest(<-conns)
	}

}

func listen(listener net.Listener, conns chan net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		log.Printf("receieve message from %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		conns <- conn
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 1024)

	for {
		size, err := conn.Read(b)
		if err != nil {
			if err == io.EOF {
				log.Println("leave.")
			} else {
				log.Printf("handle request error: %s", err)
			}
			break
		}

		conn.Write(b[:size])
	}

}
