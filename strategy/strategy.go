package strategy

import (
	"github.com/cwxstat/knight/board"
	"github.com/cwxstat/knight/knight"
)

type Record struct {
	options [][]int
	chosen  [][]int
}

func (r *Record) Add(options, chosen []int) bool {
	if len(options) > 0 {
		r.options = append(r.options, options)
		r.chosen = append(r.chosen, chosen)
		return true
	}
	return false
}

func (r *Record) Move(i []int) bool {
	return r.Add(i, i[0:1])
}

func (r *Record) Check(b *board.Board) bool {
	if len(r.options) == 0 {
		return false
	}
	i := len(r.options) - 1
	if ok := b.Move(r.options[i][0], r.options[i][1]); ok {
		return true
	}

	return false
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
