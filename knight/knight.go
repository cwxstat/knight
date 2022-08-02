package knight

import (
	"errors"
)

type Move struct {
	X, Y, Index int
}

type Knight struct {
	move    []Move
	history []int
}

func NewKnight() *Knight {
	k := &Knight{}
	k.move = append(k.move, Move{2, 1, 0})
	k.move = append(k.move, Move{1, 2, 1})

	k.move = append(k.move, Move{-2, 1, 2})
	k.move = append(k.move, Move{-1, 2, 3})

	k.move = append(k.move, Move{2, -1, 4})
	k.move = append(k.move, Move{1, -2, 5})

	k.move = append(k.move, Move{-2, -1, 6})
	k.move = append(k.move, Move{-1, -2, 7})

	return k
}

func (k *Knight) Move(index int) (Move, error) {
	if index < 0 || index >= len(k.move) {
		return Move{0, 0, 0}, errors.New("Out of bounds")
	}
	k.history = append(k.history, index)
	return k.move[index], nil

}

func (k *Knight) Next() (Move, error) {
	if len(k.history) == 0 {
		return k.Move(0)
	}
	idx := k.history[len(k.history)-1]
	k.history[idx] = (k.history[idx] % 7) + 1
	return k.Move(k.history[idx])

}

func (k *Knight) Back() int {
	if len(k.history) == 0 {
		return 0
	}
	k.history = k.history[:len(k.history)-1]
	k.move = k.move[:len(k.move)-1]
	return len(k.history)

}

func (k *Knight) Status() ([]Move, []int) {
	move := make([]Move, len(k.history))
	history := make([]int, len(k.history))
	for i, v := range k.history {
		move[i] = k.move[v]
		history[i] = v
	}
	return move, history
}

func (k *Knight) Index(i int) (Move, error) {
	if i < 0 || i >= len(k.move) {
		return Move{0, 0, 0}, errors.New("Out of bounds")
	}
	return k.move[i], nil

}

func (k *Knight) Len() int {
	return len(k.history)
}
