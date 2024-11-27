# Functional Library for Go

A lightweight library for functional programming in Go.

## Installation

```bash
go get github.com/yamada-ai/functional
```


## Example

```go
package main

import (
	"fmt"
	"github.com/yourusername/functional"
)

func main() {
	// Example usage
	result := functional.Start([]int{1, 2, 3, 4}).
		Filter(func(v int, _ int) bool {
			return v%2 == 0
		}).
		Map(func(v int, _ int) int {
			return v * v
		}).
		Result()

	fmt.Println(result) // Output: [4, 16]
}
```