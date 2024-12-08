# Functional Library for Go

A lightweight library for functional programming in Go.

## Installation

```bash
go get github.com/yamada-ai/functional
```


## Usage

### Using Functions
```go
import "github.com/yamada-ai/functional"

func main() {
    f := functional.Pipe(
        func(x int) int { return x + 1 },
        func(x int) int { return x * 2 },
    )
    println(f(5)) // 12
}
```