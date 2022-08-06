package knight_test

import (
	"fmt"
	"github.com/cwxstat/knight/knight"
)

// This is an example of knight moves
func ExampleNewKnight() {
	// Create a new knight
	k := knight.NewKnight()
	m, err := k.MoveOptions(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m.X, m.Y, m.Index)

	// Here m.X is 2, m.Y is 1, and index is 0.

	// Output:
	// 2 1 0

}
