package functional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID       int
	Name     string
	Age      int
	IsActive bool
}

func TestPipe(t *testing.T) {
	f := Pipe(
		func(x int) int { return x + 1 },
		func(x int) int { return x * 2 },
	)
	assert.Equal(t, 12, f(5)) // (5+1)*2=12
}

func TestComposeAll(t *testing.T) {
	f := ComposeAll(
		func(x int) int { return x + 1 },
		func(x int) int { return x * 2 },
	)
	assert.Equal(t, 11, f(5)) // ComposeAllは逆順適用: (5*2)+1=11
}

func TestCompose2(t *testing.T) {
	extractAge := func(u User) int { return u.Age }
	isAdult := func(age int) bool { return age >= 20 }

	f := Compose2(extractAge, isAdult)
	u := User{ID: 1, Name: "Alice", Age: 25, IsActive: true}

	assert.Equal(t, true, f(u))
}

func TestCompose3(t *testing.T) {
	u := []User{
		{ID: 1, Name: "Alice", Age: 25, IsActive: true},
		{ID: 2, Name: "Bob", Age: 30, IsActive: false},
		{ID: 3, Name: "Charlie", Age: 35, IsActive: true},
	}

	filterActive := func(users []User) []User {
		return Filter(users, func(user User) bool { return user.IsActive })
	}
	extractAges := func(users []User) []int {
		return Map(users, func(u User) int { return u.Age })
	}
	sumAges := func(ages []int) int {
		return Reduce(ages, 0, func(acc, x int) int { return acc + x })
	}

	f := Compose3(filterActive, extractAges, sumAges)
	assert.Equal(t, 60, f(u)) // 25 + 35 = 60
}

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	evens := Filter(nums, func(x int) bool { return x%2 == 0 })
	assert.Equal(t, []int{2, 4}, evens)
}

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3}
	strs := Map(nums, func(x int) string {
		return string(rune('A' + (x - 1)))
	})
	assert.Equal(t, []string{"A", "B", "C"}, strs)
}

func TestReduce(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	sum := Reduce(nums, 0, func(acc, x int) int { return acc + x })
	assert.Equal(t, 15, sum)
}
