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
	users := []User{
		{ID: 1, Name: "Alice", Age: 25, IsActive: true},
		{ID: 2, Name: "Bob", Age: 30, IsActive: false},
		{ID: 3, Name: "Charlie", Age: 35, IsActive: true},
	}

	filterActive := func(u []User) []User {
		return functional.Filter(u, func(user User) bool { return user.IsActive })
	}
	extractAges := func(u []User) []int {
		return functional.Map(u, func(user User) int { return user.Age })
	}
	sumAges := func(ages []int) int {
		return functional.Reduce(ages, 0, func(acc, x int) int { return acc + x })
	}

	// Compose2 to handle User->[]User->[]int->int is more complex, so let's do it step by step:
	process := functional.Compose3(filterActive, extractAges, sumAges)
	fmt.Println(process(users))
}
