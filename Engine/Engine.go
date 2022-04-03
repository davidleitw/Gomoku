package Engine

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	BLACKICON = "⬟"
	BLACKCODE = -1
	WHITEICON = "⬠"
	WHITECODE = 1
)

type GMKEngine struct {
	boardSize  int
	boardState [][]int

	conn *net.UnixConn
}

type Point struct {
	x, y int
}

func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

func NewEngine(boardSize int) *GMKEngine {
	log.Printf("Create Gomoku engine, board size = %d * %d\n", boardSize, boardSize)
	board := &GMKEngine{boardSize: boardSize}
	board.resetBoard()
	return board
}

func (engine *GMKEngine) resetBoard() {
	board := make([][]int, engine.boardSize)
	for i := 0; i < engine.boardSize; i++ {
		board[i] = make([]int, engine.boardSize)
	}
	engine.boardState = board
}

func (engine *GMKEngine) vaildCheckmate() ([]Point, bool) {
	var Points []Point

	return Points, false
}

// @x,@y: 棋盤座標
func (engine *GMKEngine) step(x, y int) bool {
	if x > engine.boardSize || y > engine.boardSize {
		return false
	}

	return true
}

func (engine *GMKEngine) PrintBoard() {
	engine.randomBoard()
	board := ""
	for x := 0; x < engine.boardSize; x++ {
		for y := 0; y < engine.boardSize; y++ {
			switch engine.boardState[x][y] {
			case BLACKCODE:
				board += BLACKICON
			case WHITECODE:
				board += WHITEICON
			default:
				board += "."
			}
			board += " "
		}
		board += "\n"
	}

	fmt.Println(board)
}

// To test PrintBoard
func (engine *GMKEngine) randomBoard() {
	rand.Seed(time.Now().UnixNano())
	for x := 0; x < engine.boardSize; x++ {
		for y := 0; y < engine.boardSize; y++ {
			engine.boardState[x][y] = -1 + rand.Intn(3)
		}
	}
}
