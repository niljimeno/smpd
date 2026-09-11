package main

import (
	"smpd/broadcast"
	"smpd/radio"
)

func main() {
	r := radio.NewRadio()
	broadcast.Listen(&r)
}
