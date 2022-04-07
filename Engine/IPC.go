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
	log.Printf("Create unix socket connection with %s... ", IPC_PATH)
	conn, err := net.Dial("unix", IPC_PATH)
	if err != nil {
		log.Printf("Failed!\n")
		panic(err)
	}

	log.Printf("Success! Start communication with Python RL model.\n")
	engine.conn = conn
	// buffer := make([]byte, 0, 572)
	tmp := make([]byte, 452)

	for {
		engine.sendCandiates(NewPacket(BLACKCODE, NewPoint(10, 20), NewPoint(20, 30), NewPoint(30, 40)))

		// buff := make([]byte, 87)
		// _, err = engine.conn.Read(buff)
		// if err != nil {
		// 	panic(err)
		// }
		// log.Println(buff)
		n, err := engine.conn.Read(tmp)
		if err != nil {
			panic(err)
		}

		// buffer = append(buffer, tmp[:n]...)
		log.Printf("size of data = %d, data = %d", len(tmp), tmp[:n])
		time.Sleep(2 * time.Second)
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
