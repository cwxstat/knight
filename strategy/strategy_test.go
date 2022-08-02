package strategy

import (
	"github.com/cwxstat/knight/board"
	"github.com/cwxstat/knight/knight"
	"testing"
)

func Test_Idea(t *testing.T) {
	b := board.NewBoard()
	k := knight.NewKnight()
	b.Init(8, 8)
	m, h := k.Move(0)
	t.Log(m, h)
	t.Log(k.Len())
	k.Move(1)
	t.Log(k.Status())
}
