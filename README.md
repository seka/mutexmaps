# mutexmaps

mutexmaps can be easily and thread safety for [Golang](https://golang.org/)'s map.

## Example

```
package main

import (
	"fmt"

	"github.com/seka/mutexmaps/mutexmap"
)

func main() {
	m := mutexmap.New(10)
	m.Put("key", "value")
	item := m.Get("key")
	fmt.Println(item)
}
```

