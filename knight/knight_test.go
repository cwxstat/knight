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
