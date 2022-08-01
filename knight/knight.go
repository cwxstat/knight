package knight

import (
	"sync"
)

type moves struct {
	x    int
	y    int
	hits int
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
