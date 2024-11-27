package functional

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func TestFunctionalLibrary(t *testing.T) {
	users := []User{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	}

	// テスト: Filter → Map
	result := Start(users).
		Filter(func(u User, _ int) bool {
			return u.Age >= 30
		}).
		Map(func(u User, _ int) User {
			u.Name = strings.ToUpper(u.Name)
			return u
		}).
		Result()

	// 結果の検証
	expected := []User{
		{ID: 1, Name: "ALICE", Age: 30},
		{ID: 3, Name: "CHARLIE", Age: 35},
	}
	assert.Equal(t, expected, result)
}
