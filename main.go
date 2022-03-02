package main

import (
	"fmt"

	"github.com/davidleitw/Gomoku/Engine"
)

func main() {
	engine := Engine.NewEngine(15)
	fmt.Println(engine)
}
