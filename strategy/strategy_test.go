package strategy

import (
	"testing"

	"github.com/cwxstat/knight/board"
)

func Test_Idea(t *testing.T) {
	b := board.NewBoard()
	rec := &Record{}
	b.Init(8, 8)
	b.Move(0, 0)
	r := Options(b)
	if rec.Check(b) {
		rec.Move(r)
	}
	b.Print()

}
