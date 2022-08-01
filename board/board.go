package board

import (
	"fmt"
	"sync"
)

type Moves struct {
	x, y, count int
}

type board struct {
	board   [][]int
	count   int
	history []Moves
	sync.Mutex
}

func NewBoard() *board {
	return &board{}
}

func (b *board) Init(x, y int) {
	b.Lock()
	defer b.Unlock()
	b.history = []Moves{}
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
	b.history = append(b.history, Moves{x, y, b.count})
	return true
}

func (b *board) History() []Moves {
	b.Lock()
	defer b.Unlock()
	history := []Moves{}
	for _, m := range b.history {
		history = append(history, Moves{m.x, m.y, m.count})
	}
	return history
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

func (b *board) Undo() bool {
	b.Lock()
	defer b.Unlock()
	if len(b.history) == 0 {
		return false
	}
	m := b.history[len(b.history)-1]
	b.board[m.x][m.y] = 0
	b.history = b.history[:len(b.history)-1]
	return true
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
	fmt.Println()
}
