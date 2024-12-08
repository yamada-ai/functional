package main

import (
	"fmt"

	"github.com/yamada-ai/functional"
)

type User struct {
	ID       int
	Name     string
	Age      int
	IsActive bool
}

func main() {
	// 単純なPipe例
	incDouble := functional.Pipe(
		func(x int) int { return x + 1 },
		func(x int) int { return x * 2 },
	)
	fmt.Println(incDouble(5)) // 12

	// Compose2を使った異型変換
	users := []User{
		{ID: 1, Name: "Alice", Age: 25, IsActive: true},
		{ID: 2, Name: "Bob", Age: 30, IsActive: false},
		{ID: 3, Name: "Charlie", Age: 35, IsActive: true},
	}

	filterActive := func(us []User) []User {
		return functional.Filter(us, func(u User) bool { return u.IsActive })
	}
	extractAges := func(us []User) []int {
		return functional.Map(us, func(u User) int { return u.Age })
	}
	sumAges := func(ages []int) int {
		return functional.Reduce(ages, 0, func(acc, x int) int { return acc + x })
	}

	// 2ステップ変換: []User -> []int -> int
	process := functional.Compose2(filterActive, func(users []User) int {
		return sumAges(extractAges(users))
	})

	fmt.Println(process(users)) // 60
}
