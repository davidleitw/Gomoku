package Engine

import (
	"bytes"
	"encoding/binary"
	"log"
)

const (
	GMK_PROT_HEADER = "H"
	GMK_HEADER_LEN  = 1
)

/*
封包格式:
	player_code: 1 byte  (輪到誰下棋)
	candiate_num: 1 byte (共有幾組 candiate)
	candiates: 2 * candiate_num bytes (每個候選的座標, 可以看成還沒下的部份)
*/
func NewPacket(player int, Points ...*Point) []byte {
	if player != BLACKCODE && player != WHITECODE {
		return []byte{}
	}

	buffer := bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, int8(player))
	binary.Write(buffer, binary.BigEndian, int8(len(Points)))

	for _, p := range Points {
		binary.Write(buffer, binary.BigEndian, int8(p.x))
		binary.Write(buffer, binary.BigEndian, int8(p.y))
	}
	return buffer.Bytes()
}

func ParseDicision(decision []byte) *Point {
	length := len(decision)
	if length != 2 {
		log.Println("Dicision format error!")
		return nil
	}

	x, y := int(decision[0]), int(decision[1])
	return NewPoint(x, y)
}
