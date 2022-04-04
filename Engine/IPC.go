package Engine

import (
	"log"
	"net"
	"time"
)

const (
	IPC_PATH = "./gomoku.sock"
)

func (engine *GMKEngine) buildIpcConnect() {
	conn, err := net.Dial("unix", IPC_PATH)
	if err != nil {
		panic(err)
	}

	log.Println("Create unix socket connection with ", IPC_PATH)
	engine.conn = conn
	for {
		engine.sendCandiates(NewPacket(NewPoint(10, 20), NewPoint(20, 30), NewPoint(30, 40)))

		buff := make([]byte, 87)
		_, err = engine.conn.Read(buff)
		if err != nil {
			panic(err)
		}
		log.Println(buff)
		time.Sleep(1 * time.Second)
	}
}

func (engine *GMKEngine) sendCandiates(packet []byte) {
	if engine.conn == nil {
		panic("connection fail")
	}

	_, err := engine.conn.Write(packet)
	if err != nil {
		log.Println(err)
	}
}

func (engine *GMKEngine) getDecision() []byte {
	if engine.conn == nil {
		panic("connection fail")
	}

	buf := make([]byte, 452)
	decision, err := engine.conn.Read(buf)
	if err != nil {
		log.Println(err)
	}
	return buf[0:decision]
}
