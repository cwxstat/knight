package strategy

import (
	"testing"

	"github.com/cwxstat/knight/board"
	"github.com/cwxstat/knight/knight"
)

func Test_Idea(t *testing.T) {
	b := board.NewBoard()
	b.Init(8, 8)
	b.Move(0, 0)
	r := Options(b)
	t.Log(r)

	k := knight.NewKnight()
	b.Init(8, 8)
	b.Move(0, 0)
	l, _ := b.Last()
	t.Log(l)
	m, h := k.Move(0)
	t.Log(m, h)
	t.Log(k.Len())
	k.Move(1)
	t.Log(k.Status())

	t.Log(k.MoveOptions(1))
}
