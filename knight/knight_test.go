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
	ok := b.Move(1, 2)
	if b.board[1][2].count != 1 || !ok {
		t.Errorf("Expected 1, got %d", b.board[1][2].count)
	}
	if ok := b.Move(1, 2); ok {
		t.Errorf("Expected false, got true")
	}

	b.Clear(1, 2)
	if b.board[1][2].count != 0 {
		t.Errorf("Expected 0, got %d", b.board[1][2].count)
	}

	if ok := b.Move(1, 2); !ok {
		t.Errorf("Expected true, got false")
	}
}
