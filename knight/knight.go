package knight

import (
	"fmt"
	"sync"
)

type moves struct {
	x, y, count int
}

type history struct {
	steps []moves
}

type board struct {
	board [][]int
	count int
	sync.Mutex
}

func NewBoard() *board {
	return &board{}
}

func (b *board) Init(x, y int) {
	b.Lock()
	defer b.Unlock()
	b.board = make([][]int, x)
	for i := 0; i < x; i++ {
		b.board[i] = make([]int, y)
	}
}

func (b *board) Move(x, y int) bool {
	b.Lock()
	defer b.Unlock()
	if !b.valid(x, y) {
		return false
	}
	b.count++
	b.board[x][y] = b.count
	return true
}

func (b *board) valid(x, y int) bool {
	if x < 0 || x >= len(b.board) || y < 0 || y >= len(b.board[0]) {
		return false
	}
	if b.board[x][y] != 0 {
		return false
	}
	return true
}

func (b *board) Clear(x, y int) {
	b.Lock()
	defer b.Unlock()
	b.board[x][y] = 0
}

func (b *board) Print() {
	b.Lock()
	defer b.Unlock()
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[0]); j++ {
			fmt.Printf("%d ", b.board[i][j])
		}
		fmt.Println()
	}
}
