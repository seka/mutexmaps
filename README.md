# mutexmaps

mutexmaps can be easily and thread safety for [Golang](https://golang.org/)'s map.

more information this [godoc](https://godoc.org/github.com/seka/mutexmaps).

## Example

```
package main

import (
	"fmt"

	"github.com/seka/mutexmaps/mutexmap"
	"github.com/seka/mutexmaps/mutexmultimap"
)

func main() {
	m := mutexmap.New(10)
	m.Put("key", "value")
	item := m.Get("key")
	fmt.Println(item) // "value"

    m := New(1)
    m.Put("a", "value1")
    m.Put("a", "value2")
    values := m.Get("a")
	fmt.Println(item) // []interface{}{"value1", "value2"}
}
```

## License

The MIT License (MIT)

Copyright (c) 2016 [shin sekaryo](https://github.com/seka)

