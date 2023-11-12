Put string in a box

```Go
package main

import (
	"fmt"

	"github.com/oliver-gordon/box"
)

func main() {
	fmt.Println(box.Box("I am in a box"))
	fmt.Println(box.BoxMany("I too am in a box", "a resizable box"))
}
```

```console
----------------- 
| I am in a box | 
-----------------
---------------------
| I too am in a box |
| a resizable box   |
---------------------
```