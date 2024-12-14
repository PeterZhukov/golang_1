package io

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const addr = "0.0.0.0:12345"
const proto = "tcp4"

func NetIoMain() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	var req []byte
	var buf = make([]byte, 2)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		end := false
		for _, b := range buf {
			if b == '\n' {
				end = true
				break
			}
			req = append(req, b)
		}
		if end {
			break
		}
	}

	msg := strings.TrimSuffix(string(req), "\n")
	msg = strings.TrimSuffix(msg, "\r")

	if msg == "time" {
		n, err := conn.Write([]byte(time.Now().String() + "\n"))
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(fmt.Sprintf("Клиенту отправлено %d байт", n))
	}

}
