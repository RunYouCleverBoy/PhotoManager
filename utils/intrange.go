package utils

import "fmt"

type IntRange struct {
	Min int
	Max int
}

func EmptyIntRange() *IntRange {
	return &IntRange{Min: 0, Max: -1}
}

func NewIntRange(min, max int) *IntRange {
	return &IntRange{Min: min, Max: max}
}

func (r *IntRange) Start() int {
	return r.Min
}

func (r *IntRange) End() int {
	return r.Max
}

func (r *IntRange) IsNullOrEmpty() bool {
	return r == nil || r.IsEmpty()
}

func (r *IntRange) Contains(n int) bool {
	return r.Min <= n && n <= r.Max
}

func (r *IntRange) Overlaps(other *IntRange) bool {
	return r.Contains(other.Min) || r.Contains(other.Max)
}

func (r *IntRange) IsAdjacent(other *IntRange) bool {
	return r.Min == other.Max+1 || r.Max == other.Min-1
}

func (r *IntRange) Merge(other *IntRange) *IntRange {
	return &IntRange{Min: min(r.Min, other.Min), Max: max(r.Max, other.Max)}
}

func (r *IntRange) Length() int {
	return r.Max - r.Min + 1
}

func (r *IntRange) IsEmpty() bool {
	return r.Min > r.Max
}

func (r *IntRange) String() string {
	return fmt.Sprintf("[%d, %d]", r.Min, r.Max)
}

func (r *IntRange) Equals(other *IntRange) bool {
	return r.Min == other.Min && r.Max == other.Max
}

func (r *IntRange) Clone() *IntRange {
	return &IntRange{Min: r.Min, Max: r.Max}
}

func (r *IntRange) ForEach(f func(int)) {
	for i := r.Min; i <= r.Max; i++ {
		f(i)
	}
}
