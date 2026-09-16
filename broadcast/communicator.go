package broadcast

import (
	"fmt"
	"net"
	"smpd/radio"
	"strings"
)

func processCommand(data []byte) []string {
	result := []string{}
	current := strings.Builder{}

	literal := false
	inBlock := false
	var blockMethod byte

	applyChanges := func() {
		newAddition := current.String()
		if newAddition == "" {
			return
		}

		result = append(result, newAddition)
		current = strings.Builder{}
	}

	for _, b := range data {
		switch {
		default:
			current.WriteByte(b)

		case literal:
			current.WriteByte(b)
			literal = false

		case b == 0:
			continue

		case b == '\\':
			literal = true

		case inBlock && b == blockMethod:
			inBlock = false

		case inBlock:
			current.WriteByte(b)

		case b == ' ', b == '\n':
			applyChanges()

		case b == '"', b == '\'', b == '`':
			inBlock = true
			blockMethod = b
		}
	}

	applyChanges()
	fmt.Println(result)
	return result
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
