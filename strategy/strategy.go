package strategy

import (
	"github.com/cwxstat/knight/board"
	"github.com/cwxstat/knight/knight"
)

type Record struct {
	options [][]int
	chosen  [][]int
}

func (r *Record) Add(options, chosen []int) {
	r.options = append(r.options, options)
	r.chosen = append(r.chosen, chosen)
}

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
