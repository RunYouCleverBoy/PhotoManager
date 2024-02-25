package utils

import "fmt"

type IntRange[T int | int8 | int16 | int32 | int64] struct {
	Min T
	Max T
}

func EmptyIntRange[T int | int8 | int16 | int32 | int64]() *IntRange[T] {
	return &IntRange[T]{Min: 0, Max: -1}
}

func NewIntRange[T int | int8 | int16 | int32 | int64](min, max T) *IntRange[T] {
	return &IntRange[T]{Min: min, Max: max}
}

func (r *IntRange[T]) Start() T {
	return r.Min
}

func (r *IntRange[T]) End() T {
	return r.Max
}

func (r *IntRange[T]) IsNullOrEmpty() bool {
	return r == nil || r.IsEmpty()
}

func (r *IntRange[T]) Contains(n T) bool {
	return r.Min <= n && n <= r.Max
}

func (r *IntRange[T]) Overlaps(other *IntRange[T]) bool {
	return r.Contains(other.Min) || r.Contains(other.Max)
}

func (r *IntRange[T]) IsAdjacent(other *IntRange[T]) bool {
	return r.Min == other.Max+1 || r.Max == other.Min-1
}

func (r *IntRange[T]) Merge(other *IntRange[T]) *IntRange[T] {
	return &IntRange[T]{Min: min(r.Min, other.Min), Max: max(r.Max, other.Max)}
}

func (r *IntRange[T]) Length() T {
	return r.Max - r.Min + 1
}

func (r *IntRange[T]) IsEmpty() bool {
	return r.Min > r.Max
}

func (r *IntRange[T]) String() string {
	return fmt.Sprintf("[%d, %d]", r.Min, r.Max)
}

func (r *IntRange[T]) Equals(other *IntRange[T]) bool {
	return r.Min == other.Min && r.Max == other.Max
}

func (r *IntRange[T]) Clone() *IntRange[T] {
	return &IntRange[T]{Min: r.Min, Max: r.Max}
}

func (r *IntRange[T]) ForEach(f func(T)) {
	for i := r.Min; i <= r.Max; i++ {
		f(i)
	}
}
