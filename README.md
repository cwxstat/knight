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
	k.Move(0)
	k.Move(1)
	m, h := k.Status()
	fmt.Println(h)
	fmt.Println(m)
}

```


```bash




git tag -a v0.0.1 -m "Alpha release"

```