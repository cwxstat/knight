# knight

[Playground](https://go.dev/play/p/aA6BOyiVIoh)
```go
package main

import (
	"fmt"

	"github.com/cwxstat/knight/knight"
)

func main() {
	k := knight.NewKnight()
	m2,_ :=k.Move(0)
    fmt.Println(m2.X)
	k.Move(1)
	m, h := k.Status()
	fmt.Println(h)
	fmt.Println(m)
}

```


```bash



git tag -a v0.0.2 -m "Alpha release"


```