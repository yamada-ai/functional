package functional

import (
	"github.com/samber/lo"
)

// Pipe構造体: チェーン可能な構造
type Pipe[T any] struct {
	items []T
}

// チェーン開始
func Start[T any](items []T) Pipe[T] {
	return Pipe[T]{items}
}

// Filterメソッド: 条件に合う要素を抽出
func (p Pipe[T]) Filter(predicate func(T, int) bool) Pipe[T] {
	return Pipe[T]{lo.Filter(p.items, predicate)}
}

// Mapメソッド: 要素を変換
func (p Pipe[T]) Map(mapper func(T, int) T) Pipe[T] {
	return Pipe[T]{lo.Map(p.items, mapper)}
}

// 結果を取得
func (p Pipe[T]) Result() []T {
	return p.items
}
