package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

func main() {

	// Assigns the incoming traffic protocol and port
	addr, err := net.ResolveUDPAddr("udp", ":9999")

	if err != nil {
		log.Fatal(err)
	}
	// Listens incoming UDP packets
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Listening on :9999")

	// Buffer which holds incoming packets
	buf := make([]byte, 324)

	// Loops forever and reads one packet at a time
	for {
		// if there is a error or packet is smaller than 324 it ignores it and wait for another
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil || n < 324 {
			continue
		}

		var t Telemetry

		// Takes packet bytes and makes them readable
		reader := bytes.NewReader(buf[:n])

		//Takes raw bytes and fills the telemetry struct
		if err := binary.Read(reader, binary.LittleEndian, &t); err != nil {
			log.Println("decode error", err)
			continue
		}

		speedKmh := t.Speed * 3.6
		fmt.Printf("\rRaceon: %d | Speed: %.1f km/h | RPM: %.0f | Gear: %d  ", t.IsRaceOn, speedKmh, t.CurrentEngineRpm, t.Gear)
	}
}
