package strategy

import (
	"github.com/cwxstat/knight/board"
	"github.com/cwxstat/knight/knight"
)

func Do() {
	b := board.NewBoard()
	k := knight.NewKnight()
	b.Init(8, 8)
	k.Move(0)
}
