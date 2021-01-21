package main

import (
	"flag"
	"log"
	"quad"
)

var (
	port string
	baud int
)

func init() {
	flag.StringVar(&port, "p", "/dev/ttyAMA0", "serial port device")
	flag.IntVar(&baud, "b", 9600, "baud rate")
	flag.Parse()
}

func main() {
	lPort, err := ltm.Make(port, baud)
	if err != nil {
		log.Fatalln(err)
	}

	go lPort.Read()

	for {
		lat, lon := lPort.GetGPS()
		log.Printf("lat: %d, lon: %d", lat, lon)
		bat := lPort.GetBat()
		log.Printf("Bat: %d", bat)
	}
}