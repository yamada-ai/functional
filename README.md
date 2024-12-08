# Functional Library for Go

A lightweight library for functional programming in Go.

## Installation

```bash
go get github.com/yamada-ai/functional
```

## Quick Start
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

## Basic Usage
### Pipe (Same-Type Transformations)
- Pipe applies a series of func(T) T transformations in the order they are listed. It’s intuitive and straightforward.

```go
f := functional.Pipe(
    func(x int) int { return x + 10 },
    func(x int) int { return x * 3 },
)
fmt.Println(f(2)) // ((2+10)*3) = 36
```

### ComposeAll (Reverse Application Order)
- ComposeAll applies functions in reverse order (last listed function is applied first).

```go
g := functional.ComposeAll(
    func(x int) int { return x + 10 },
    func(x int) int { return x * 3 },
)
fmt.Println(g(2)) // (2*3)+10 = 16
```

## Slightly More Advanced: Changing Types
- For transformations across different types, use Compose2 or Compose3:

### Compose2 (A->B->C)
```go
extractLength := func(s string) int { return len(s) }
isEven := func(n int) bool { return n%2 == 0 }

checkEvenLength := functional.Compose2(extractLength, isEven)
fmt.Println(checkEvenLength("Hello")) // len("Hello")=5 -> 5%2!=0 => false
```

### Compose3 (A->B->C->D)
```go
toRuneSlice := func(s string) []rune { return []rune(s) }
toInts := func(runes []rune) []int {
    return functional.Map(runes, func(r rune) int { return int(r) })
}
sumInts := func(nums []int) int {
    return functional.Reduce(nums,0,func(acc,x int)int{return acc+x})
}

process := functional.Compose3(toRuneSlice, toInts, sumInts)
fmt.Println(process("ABC")) // 'A'=65,'B'=66,'C'=67 => sum=198
```

## Data Processing with Filter, Map, Reduce
### Filter & Map
```go
nums := []int{1, 2, 3, 4, 5}
evens := functional.Filter(nums, func(x int) bool {return x%2==0}) // [2,4]
doubled := functional.Map(evens, func(x int) int {return x*2})     // [4,8]
fmt.Println(doubled)
```
### Reduce
```go
nums := []int{1, 2, 3, 4, 5}
sum := functional.Reduce(nums,0, func(acc,x int) int {return acc+x})
fmt.Println(sum) // 15
```

## More Complex Examples
### Integrating Filter/Map/Reduce with Pipe
```go
// We want to take active users, extract their ages, and compute the total
type User struct {
    ID       int
    Name     string
    Age      int
    IsActive bool
}

users := []User{
    {ID:1, Name:"Alice", Age:25, IsActive:true},
    {ID:2, Name:"Bob", Age:30, IsActive:false},
    {ID:3, Name:"Charlie", Age:35, IsActive:true},
}

filterActive := func(u []User) []User {
    return functional.Filter(u, func(user User) bool {return user.IsActive})
}
extractAges := func(u []User) []int {
    return functional.Map(u, func(user User) int {return user.Age})
}
sumAges := func(ages []int) int {
    return functional.Reduce(ages,0, func(acc,x int)int{return acc+x})
}

// Compose2 to handle User->[]User->[]int->int is more complex, so let's do it step by step:
activeUsers := filterActive(users)    // [Alice(25), Charlie(35)]
ages := extractAges(activeUsers)      // [25, 35]
total := sumAges(ages)                // 60
fmt.Println(total)
```
- If you want a single call, Compose3 can help:

```go
process := functional.Compose3(filterActive, extractAges, sumAges)
fmt.Println(process(users)) // 60
```

### Grouping Data (GroupBy)
- GroupBy lets you categorize data by a key function:
```go
grouped := functional.GroupBy(users, func(u User) bool { return u.IsActive })
// map[bool][]User
// true  => [Alice(25), Charlie(35)]
// false => [Bob(30)]

for k, grp := range grouped {
    fmt.Println("Key:", k, "Users:", grp)
}
```
- Another example: grouping strings by their first letter:

```go
words := []string{"apple", "banana", "apricot", "blueberry", "cherry"}
byFirstChar := functional.GroupBy(words, func(w string) rune {
    return []rune(w)[0]
})
// 'a' => ["apple","apricot"]
// 'b' => ["banana","blueberry"]
// 'c' => ["cherry"]
```