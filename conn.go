package ws

const (
	finalBit = 1 << 7
	rsv1Bit  = 1 << 6
	rsv2Bit  = 1 << 5
	rsv3Bit  = 1 << 4

	maskBit                    = 1 << 7
	maxFrameHeaderSize         = 2 + 8 + 4 // Fixed header + length + mask
	maxControlFramePayloadSize = 125

	continuationFrame    = 0x00
	textFrame            = 0x01
	binaryFrame          = 0x02
	connectionCloseFrame = 0x08
	pingFrame            = 0x09
	pongFrame            = 0xA

	noFrame = -1
)

type Conn struct {
}

func newConn() {

}
