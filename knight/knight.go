package knight

import (
	"sync"
)

type moves struct {
	x     int
	y     int
	count int
}

type board struct {
	board [][]moves
	sync.Mutex
}

func NewBoard() *board {
	return &board{}
}

func (b *board) Init(x, y int) {
	b.Lock()
	defer b.Unlock()
	b.board = make([][]moves, x)
	for i := 0; i < x; i++ {
		b.board[i] = make([]moves, y)
	}
}

func (b *board) Move(x, y int) {
	b.Lock()
	defer b.Unlock()
	b.board[x][y].count++
}

func (b *board) Clear(x, y int) {
	b.Lock()
	defer b.Unlock()
	b.board[x][y].count = 0
}
