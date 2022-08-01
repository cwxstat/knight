package knight

type Move struct {
	x, y, index int
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
		return Move{0, 0, 0}, nil
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
