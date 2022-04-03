package main

import (
	"github.com/davidleitw/Gomoku/Engine"
)

func main() {
	// engine := Engine.NewEngine(15)
	// engine.PrintBoard()
	p1 := Engine.NewPoint(24, 61)
	p2 := Engine.NewPoint(36, 199)
	data := Engine.NewPacket(p1, p2)
	Engine.Unpack(data)
}
