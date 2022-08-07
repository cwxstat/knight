[![Go](https://github.com/cwxstat/knight/actions/workflows/go.yml/badge.svg)](https://github.com/cwxstat/knight/actions/workflows/go.yml)
# knight

[Playground](https://go.dev/play/p/_N5NER84Ff8)
```go
-- go.mod --
module dev

go 1.18

require (
	github.com/cwxstat/knight v0.0.6
)
-- main.go --
package main

import (
	"fmt"

	"github.com/cwxstat/knight/board"
	"github.com/cwxstat/knight/knight"
)

func main() {
	k := knight.NewKnight()
	b := board.NewBoard()
	b.Init(8, 8)
	m2, _ := k.Move(0)
	fmt.Println(m2.X, "H")
	k.Move(1)
	m, h := k.Status()
	fmt.Println(h)
	fmt.Println(m)
}


```


```bash
https://pkg.go.dev/github.com/cwxstat/knight@v0.0.6/knight#example-NewKnight

Ref: https://github.com/gonum/gonum

https://go.dev/blog/examples

go install -v golang.org/x/tools/cmd/godoc@latest


git tag -a v0.0.6 -m "Documentation test"


```
