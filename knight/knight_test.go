package knight

import (
	"reflect"
	"testing"
)

func TestKnight_Move(t *testing.T) {
	k := NewKnight()
	move, err := k.Move(2)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(move, Move{-2, 1, 2}) {
		t.Errorf("Expected %v, got %v", Move{-2, 1, 2}, move)
	}
}

func TestKnight_Next(t *testing.T) {
	k := NewKnight()

	for i := 0; i < 8; i++ {
		move, _ := k.Next()
		if !reflect.DeepEqual(move, k.move[i]) {
			t.Errorf("Expected %v, got %v", k.move[i], move)
		}
	}

}

func TestKnight_Back(t *testing.T) {
	k := NewKnight()

	k.Move(2)
	k.Move(1)
	k.Move(1)
	move, history := k.Status()
	t.Log(move, history)
	k.Back()
	move, history = k.Status()
	if len(move) != 2 {
		t.Errorf("Expected %v, got %v", 2, len(move))
	}
	k.Back()
	k.Back()
	k.Back()
	move, history = k.Status()
	if len(move) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(move))
	}
}

func TestKnight_X(t *testing.T) {
	k := NewKnight()

	m, _ := k.Move(2)
	t.Log(m.X)
}
