package Engine

import (
	"fmt"
	"log"
	"net"
)

const (
	BLACKICON = "⬟"
	BLACKCODE = 1
	WHITEICON = "⬠"
	WHITECODE = 2
)

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	if x > 255 || y > 255 {
		log.Println("x, y must smaller than 255")
		return nil
	}
	return &Point{x: x, y: y}
}

type GMKEngine struct {
	boardSize  int
	boardState [][]int
	player     int
	conn       net.Conn
}

func NewEngine(boardSize int) *GMKEngine {
	log.Printf("Create Gomoku engine, board size = %d * %d\n", boardSize, boardSize)
	engine := &GMKEngine{boardSize: boardSize, player: BLACKCODE}
	return engine
}

func (engine *GMKEngine) Run() {
	engine.resetBoard()
	engine.buildIpcConnect()
	engine.printBoard()
	for {
		candiates := engine.allPossibleCandiates()
		engine.sendCandiates(NewPacket(engine.Player(), candiates...))

		// 接收 RL 模型選擇要下的位置， 並根據回傳結果更新棋盤狀態
		var decision *Point = ParseDicision(engine.getDecision())
		engine.step(decision)

		engine.printBoard() // For demo
		engine.changePlayer()
	}
}

func (engine *GMKEngine) resetBoard() {
	board := make([][]int, engine.boardSize)
	for i := 0; i < engine.boardSize; i++ {
		board[i] = make([]int, engine.boardSize)
	}
	engine.boardState = board
}

func (engine *GMKEngine) allPossibleCandiates() []*Point {
	ps := make([]*Point, 0)
	for x := 0; x < engine.boardSize; x++ {
		for y := 0; y < engine.boardSize; y++ {
			if engine.boardState[x][y] == 0 {
				ps = append(ps, NewPoint(x, y))
			}
		}
	}
	return ps
}

func (engine *GMKEngine) Player() int {
	return engine.player
}

func (engine *GMKEngine) changePlayer() {
	engine.player = engine.player%2 + 1
}

func (engine *GMKEngine) step(p *Point) {
	engine.boardState[p.x][p.y] = engine.Player()
}

func (engine *GMKEngine) printBoard() {
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
