package Engine

import (
	"bytes"
	"encoding/binary"
	"log"
)

const (
	GMK_PROT_HEADER = "Headers"
	GMK_HEADER_LEN  = 7
)

func NewPacket(Points ...Point) []byte {
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write([]byte(GMK_PROT_HEADER))
	binary.Write(buffer, binary.BigEndian, int8(len(Points)))

	for _, p := range Points {
		binary.Write(buffer, binary.BigEndian, int8(p.x))
		binary.Write(buffer, binary.BigEndian, int8(p.y))
	}
	return buffer.Bytes()
}

func Unpack(packet []byte) {
	length := len(packet)
	if length < GMK_HEADER_LEN || string(packet[0:GMK_HEADER_LEN]) != "Headers" {
		return
	}

	num := int8(packet[GMK_HEADER_LEN])
	if length != GMK_HEADER_LEN+1+int(num*2) {
		return
	}

	points := make([]Point, 0)
	idx := GMK_HEADER_LEN + 1
	for idx != length {
		x := int(packet[idx])
		idx++
		y := int(packet[idx])
		idx++
		points = append(points, NewPoint(x, y))
	}
	log.Println(points)
}
