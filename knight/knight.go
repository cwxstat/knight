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

func (b *board) Move(x, y int) bool {
	b.Lock()
	defer b.Unlock()
	if !b.valid(x, y) {
		return false
	}
	b.board[x][y].count++
	return true
}

func (b *board) valid(x, y int) bool {
	if x < 0 || x >= len(b.board) || y < 0 || y >= len(b.board[0]) {
		return false
	}
	if b.board[x][y].count != 0 {
		return false
	}
	return true
}

func (b *board) Clear(x, y int) {
	b.Lock()
	defer b.Unlock()
	b.board[x][y].count = 0
}
