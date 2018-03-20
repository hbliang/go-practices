package main

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"time"
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

	conn.SetDeadline(time.Now().Add(10 * time.Second))

	for {
		size, err := conn.Read(b)

		if err != nil {
			if err == io.EOF {
				log.Println("leave.")
			} else {
				log.Println("handle request error: ", err)
			}
			break
		}

		log.Println("request data: \n", string(b[:size]))

		req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(b)))
		if err != nil {
			log.Println("make request errors: ", err)
			return
		}

		var serverHost string
		if req.URL.Port() == "" {
			serverHost = req.URL.Host + ":80"
		} else {
			serverHost = req.URL.Host
		}

		server, err := net.Dial("tcp", serverHost)

		if err != nil {
			log.Println("dial server error: ", err)
			return
		}
		defer server.Close()

		if req.Method == http.MethodConnect {
			conn.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
			go io.Copy(server, conn)
		} else {
			server.Write(b[:size])
		}

		io.Copy(conn, server)
	}
}
