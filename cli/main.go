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

	go lPort.Read()

	for {
		lat, lon := lPort.GetGPS()
		log.Printf("lat: %d, lon: %d", lat, lon)
	}
}