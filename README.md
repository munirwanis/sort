# sort
Sort library made with go

This library was made to implement the various sorts we made when we're studying algorithms.

The first sort implemented here is BubbleSort

## Using

You can use this code as an example:
```go
package main

import (
	"fmt"
	"math/rand"

	"github.com/munirwanis/sort"
)

func main() {
	var array = rand.Perm(10) // Creates random array with 10 items in it
  
	fmt.Println(array)
	sort.BubbleSort(array)
	fmt.Println(array)
}

```
