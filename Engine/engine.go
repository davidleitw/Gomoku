package Engine

import "log"

type GMKEngine struct {
	boardSize int
}

func NewEngine(boardSize int) *GMKEngine {
	log.Printf("Create Gomoku engine, board size = %d * %d\n", boardSize, boardSize)
	return &GMKEngine{boardSize: boardSize}
}
