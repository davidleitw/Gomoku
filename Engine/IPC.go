package Engine

import (
	"log"
	"net"
)

const (
	IPC_PATH = "/var/run/gomoku.sock"
)

func (engine *GMKEngine) BuildIpcConnect() {
	sockaddr, err := net.ResolveUnixAddr("unix", IPC_PATH)
	if err != nil {
		log.Println(err)
		return
	}

	sockconn, err := net.DialUnix("unix", nil, sockaddr)
	if err != nil {
		log.Println(err)
		return
	}
	engine.conn = sockconn
	engine.boardStateRoutine()
	defer engine.conn.Close()
}

func (engine *GMKEngine) boardStateRoutine() {
	if engine.conn == nil {
		return
	}
	go onMessageReceived(engine.conn)
}

func onMessageReceived(conn *net.UnixConn) {

}

func (engine *GMKEngine) sendState(state []byte) {
	if engine.conn == nil {
		return
	}

}
