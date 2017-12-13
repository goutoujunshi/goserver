package cons

import (
	"net"
	"fmt"
	"time"
)

func Server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	for {
		b := make([]byte, 1024)
		conn.Read(b)
		fmt.Println(string(b))
	}

	conn.Close()
}

func Client() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	for {
		i++
		fmt.Fprintf(conn,"xiehongbo%d\n",i)

		time.Sleep(time.Second * 2)
	}

	conn.Close()
}




