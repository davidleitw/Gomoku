package Engine

import (
	"log"
	"net"
)

const (
	IPC_PATH = "./gomoku.sock"
)

func (engine *ReversiEngine) BuildIpcConnect() {
	log.Printf("Create unix socket connection with %s... ", IPC_PATH)
	conn, err := net.Dial("unix", IPC_PATH)
	if err != nil {
		log.Printf("Failed!\n")
		panic(err)
	}

	log.Printf("Success! Start communication with Python RL model.\n")
	engine.conn = conn
}

func (engine *ReversiEngine) SendCandiates(packet []byte) {
	if engine.conn == nil {
		panic("connection fail")
	}

	_, err := engine.conn.Write(packet)
	if err != nil {
		log.Println(err)
	}
}

func (engine *ReversiEngine) GetDecision() []byte {
	if engine.conn == nil {
		panic("connection fail")
	}

	buf := make([]byte, 2)
	_, err := engine.conn.Read(buf)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return buf
}
