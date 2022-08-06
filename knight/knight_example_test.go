package knight_test

import (
	"fmt"
	"github.com/cwxstat/knight/knight"
)

func ExampleMoves() {
	k := knight.NewKnight()
	m, err := k.MoveOptions(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)

	// Here m.X is 2, m.Y is 1, and index is 0.

	// Output:
	// {2 1 0}

}
