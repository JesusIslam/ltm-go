package main

import (
	"log"
	"quad"
)

func main() {
	lPort, err := ltm.Make("/dev/ttyUSB0", 9600)
	if err != nil {
		log.Fatalln(err)
	}

	lPort.Read()
}