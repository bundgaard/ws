package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

func fail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	parseFrame([]byte{1, 4})
	os.Exit(1)

	remoteAddr, err := net.ResolveTCPAddr("tcp", ":8888")
	check(err)
	conn, err := net.DialTCP("tcp", nil, remoteAddr)
	fail(err)
	defer conn.Close()
	for {
		time.Sleep(1 * time.Second)
	}
}

func parseFrame(b []byte) {

	// Parse Header 2 byte 16 bits
	firstByte := uint8(b[0])
	final := (firstByte >> 7) & 0x1
	reserved1, reserved2, reserved3 := ((firstByte >> 6) & 0x1), firstByte>>5&0x1, firstByte>>4&0x1
	opCode := firstByte & 0xf

	fmt.Println(final, reserved1, reserved2, reserved3, opCode)
}
