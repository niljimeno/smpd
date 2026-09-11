package broadcast

import (
	"net"
	"smpd/radio"
	"strings"
)

func processCommand(data []byte) []string {
	return strings.Fields(string(data))
}

func dial(conn net.Conn, r *radio.Radio) {
	defer conn.Close()

	for {
		data := make([]byte, 512)
		_, err := conn.Read(data)

		if err != nil {
			return
		}

		command := processCommand(data)
		output := r.Execute(command)
		if output == "" {
			continue
		}

		conn.Write([]byte(output + "\n"))
	}
}

func Listen(r *radio.Radio) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go dial(conn, r)
	}
}
