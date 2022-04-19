package Engine

import (
	"fmt"
	"log"
	"net"
	"sync"
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
	return &Point{x: x, y: y}
}

type board struct {
	size      int
	blackCnt  int
	whiteCnt  int
	player    int
	state     [][]int
	remaining int
}

func NewBoard(size int) *board {
	brd := &board{size: size, blackCnt: 0, whiteCnt: 0, player: BLACKCODE}
	state := make([][]int, size)
	for i := 0; i < size; i++ {
		state[i] = make([]int, size)
	}
	brd.state = state
	brd.remaining = size * size
	return brd
}

func (brd *board) reset() {
	for x := 0; x < brd.size; x++ {
		for y := 0; y < brd.size; y++ {
			brd.state[x][y] = 0
		}
	}
	brd.blackCnt = 0
	brd.whiteCnt = 0
	brd.player = BLACKCODE
	brd.remaining = brd.size * brd.size
}

func (brd *board) Player() int {
	return brd.player
}

func (brd *board) changePlayer() {
	brd.player = brd.player%2 + 1
}

func (brd *board) AllPossibleCandiates() []*Point {
	fmt.Println(brd.remaining)
	points := make([]*Point, brd.remaining)
	idx := 0
	for x := 0; x < brd.size; x++ {
		for y := 0; y < brd.size; y++ {
			if brd.state[x][y] == 0 {
				points[idx] = NewPoint(x, y)
				idx++
			}
		}
	}
	return points
}

func (brd *board) printBoard() {
	b := ""
	for x := 0; x < brd.size; x++ {
		for y := 0; y < brd.size; y++ {
			switch brd.state[x][y] {
			case BLACKCODE:
				b += BLACKICON
			case WHITECODE:
				b += WHITEICON
			default:
				b += "."
			}
			b += " "
		}
		if x != brd.size-1 {
			b += "\n"
		}
	}
	fmt.Println(b)
	fmt.Println()
	fmt.Println()
}

// 遊戲主要邏輯
func (brd *board) step(p *Point) {
	// 下一步, 判斷有沒有要翻面
	if brd.state[p.x][p.y] == 0 {
		brd.state[p.x][p.y] = brd.Player()
	} else {
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(8)
	go brd.judge(p.x, p.y, -1, 0, wg)
	go brd.judge(p.x, p.y, -1, -1, wg)
	go brd.judge(p.x, p.y, 0, -1, wg)
	go brd.judge(p.x, p.y, 1, -1, wg)
	go brd.judge(p.x, p.y, 1, 0, wg)
	go brd.judge(p.x, p.y, 1, 1, wg)
	go brd.judge(p.x, p.y, 0, 1, wg)
	go brd.judge(p.x, p.y, -1, 1, wg)
	wg.Wait()

	// 更新 black, white count, 換對手下
	brd.updateCount()
	brd.changePlayer()
	brd.remaining--
}

func (brd *board) judge(targetX, targetY int, dirX, dirY int, wg *sync.WaitGroup) {
	defer wg.Done()
	stack := make([]Point, 0)

	for {
		targetX += dirX
		targetY += dirY

		if targetX < 0 || targetX > brd.size-1 || targetY < 0 || targetY > brd.size-1 || brd.state[targetX][targetY] == 0 {
			return
		}

		if brd.state[targetX][targetY] == brd.Player() {
			break
		}

		stack = append(stack, Point{x: targetY, y: targetY})
	}

	for _, p := range stack {
		brd.state[p.x][p.y] = brd.Player()
	}
}

func (brd *board) updateCount() {
	black := 0
	white := 0
	for x := 0; x < brd.size; x++ {
		for y := 0; y < brd.size; y++ {
			switch brd.state[x][y] {
			case BLACKCODE:
				black++
			case WHITECODE:
				white++
			default:
			}
		}
	}
	brd.blackCnt = black
	brd.whiteCnt = white
}

func (brd *board) GameOver() bool {
	return brd.remaining == 0
}

type ReversiEngine struct {
	Epochs          int
	BoardSize       int
	Board           *board
	conn            net.Conn
	EngineDebugMode bool
}

func NewEngine(BoardSize, epochs int) *ReversiEngine {
	log.Printf("Create Reversi engine, board size = %d * %d\n", BoardSize, BoardSize)
	engine := &ReversiEngine{Epochs: epochs, BoardSize: BoardSize, EngineDebugMode: true}
	return engine
}

func (engine *ReversiEngine) Run() {
	engine.BuildIpcConnect()
	for epoch := 0; epoch < engine.Epochs; epoch++ {
		engine.resetBoard()
		engine.PrintBoard()
		for {
			candiates := engine.Board.AllPossibleCandiates()
			engine.SendCandiates(NewPacket(engine.Board.Player(), candiates...))

			// 接收 RL 模型選擇要下的位置， 並根據回傳結果更新棋盤狀態
			var decision *Point = ParseDicision(engine.GetDecision())
			engine.Step(decision)

			if engine.Board.GameOver() {
				// 回傳遊戲結束給 RL 模型
				break
			}
			engine.PrintBoard()
		}
	}
}

func (engine *ReversiEngine) resetBoard() {
	if engine.Board == nil {
		engine.Board = NewBoard(engine.BoardSize)
	} else {
		engine.Board.reset()
	}
}

func (engine *ReversiEngine) Step(p *Point) {
	engine.Board.step(p)
}

func (engine *ReversiEngine) PrintBoard() {
	if engine.EngineDebugMode {
		engine.Board.printBoard()
	}
}

func (engine *ReversiEngine) Checkmate(step *Point) bool {
	x, y := step.x, step.y
	if x < 0 || x > engine.BoardSize-1 || y < 0 || y > engine.BoardSize-1 {
		return false
	}
	return false
}
