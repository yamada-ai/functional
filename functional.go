package functional

// Filter: スライスを条件でフィルタリング
func Filter[T any](s []T, predicate func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range s {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

// Map: スライス要素を別の型に変換
func Map[T any, R any](s []T, mapper func(T) R) []R {
	mapped := make([]R, len(s))
	for i, v := range s {
		mapped[i] = mapper(v)
	}
	return mapped
}

// Reduce: スライス要素を累積処理
func Reduce[T any, R any](s []T, initial R, reducer func(R, T) R) R {
	acc := initial
	for _, v := range s {
		acc = reducer(acc, v)
	}
	return acc
}

// GroupBy: キーに応じて分類
func GroupBy[T any, K comparable](s []T, keyFunc func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range s {
		key := keyFunc(v)
		result[key] = append(result[key], v)
	}
	return result
}

// 基本的なパイプライン
func Pipe[T any](funcs ...func(T) T) func(T) T {
	return func(arg T) T {
		for _, f := range funcs {
			arg = f(arg)
		}
		return arg
	}
}

// 同様にComposeAll (逆順) を用意してもよい:
func ComposeAll[T any](funcs ...func(T) T) func(T) T {
	return func(arg T) T {
		for i := len(funcs) - 1; i >= 0; i-- {
			arg = funcs[i](arg)
		}
		return arg
	}
}

// 2ステップ型変換用Compose
func Compose2[A, B, C any](f1 func(A) B, f2 func(B) C) func(A) C {
	return func(arg A) C {
		return f2(f1(arg))
	}
}

// 3ステップ版（必要なら）
func Compose3[A, B, C, D any](f1 func(A) B, f2 func(B) C, f3 func(C) D) func(A) D {
	return func(arg A) D {
		return f3(f2(f1(arg)))
	}
}
