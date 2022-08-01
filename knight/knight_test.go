package knight

import "testing"

func TestBoard_Init(t *testing.T) {
	b := NewBoard()
	b.Init(5, 5)
	if len(b.board) != 5 {
		t.Errorf("Expected 5 rows, got %d", len(b.board))
	}
	if len(b.board[0]) != 5 {
		t.Errorf("Expected 5 columns, got %d", len(b.board[0]))
	}
}

func Test_Move(t *testing.T) {
	b := NewBoard()
	b.Init(5, 5)
	b.Move(1, 2)
	if b.board[1][2].count != 1 {
		t.Errorf("Expected 1, got %d", b.board[1][2].count)
	}
	b.Clear(1, 2)
	if b.board[1][2].count != 0 {
		t.Errorf("Expected 0, got %d", b.board[1][2].count)
	}
}
