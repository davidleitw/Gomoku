package main

import (
	"github.com/davidleitw/Gomoku/Engine"
)

func main() {
	engine := Engine.NewEngine(15, 100)
	engine.Run()
}
