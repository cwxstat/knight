package strategy

import (
	"github.com/cwxstat/knight/board"
	"github.com/cwxstat/knight/knight"
)

func Options(b *board.Board) []int {
	valid := []int{}
	k := knight.NewKnight()
	for i := 0; i < 8; i++ {
		m, err := k.MoveOptions(i)
		if err != nil {
			continue
		}
		if ok := b.ValidJump(m.X, m.Y); ok {
			valid = append(valid, i)
		}
	}
	return valid
}
