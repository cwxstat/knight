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
	if b.board[1][2] != 1 || !ok {
		t.Errorf("Expected 1, got %d", b.board[1][2])
	}
	if ok := b.Move(1, 2); ok {
		t.Errorf("Expected false, got true")
	}

	b.Undo()
	b.Undo()
	if b.board[1][2] != 0 {
		t.Errorf("Expected 0, got %d", b.board[1][2])
	}

	if ok := b.Move(1, 2); !ok {
		t.Errorf("Expected true, got false")
	}
}

func Test_Print(t *testing.T) {
	b := NewBoard()
	b.Init(8, 8)
	b.Move(0, 0)
	b.Move(0, 1)
	b.Move(0, 2)
	b.Print()
	b.Undo()
	b.Print()
}

func Test_History(t *testing.T) {
	b := NewBoard()
	b.Init(8, 8)
	b.Move(0, 0)
	b.Move(0, 1)
	b.Move(0, 2)
	h := b.History()
	if len(h) != 3 {
		t.Errorf("Expected 3 moves, got %d", len(h))
	}
	b.Undo()
	h = b.History()
	if len(h) != 2 {
		t.Errorf("Expected 2 moves, got %d", len(h))
	}
}
