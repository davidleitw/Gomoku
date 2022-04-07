package main

import (
	"github.com/davidleitw/Gomoku/Engine"
)

func main() {
	engine := Engine.NewEngine(15)
	engine.Run()
	engine.PrintBoard()
	// p1 := Engine.NewPoint(24, 61)
	// p2 := Engine.NewPoint(36, 199)
	// p3 := Engine.NewPoint(87, 146)
	// p4 := Engine.NewPoint(255, 78)
	// data := Engine.NewPacket(p1, p2, p3, p4)
	// Engine.Unpack(data)
}
