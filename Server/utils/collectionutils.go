package utils

import "strings"

func Map[T any, U any](arr *[]T, f func(*T) U) []U {
	result := make([]U, len(*arr))
	for i, v := range *arr {
		result[i] = f(&v)
	}
	return result
}

func MapOrFail[T any, U any](arr *[]T, f func(T) (U, error)) (*[]U, error) {
	result := make([]U, len(*arr))
	for i, v := range *arr {
		u, err := f(v)
		if err != nil {
			return nil, err
		}
		result[i] = u
	}
	return &result, nil
}

func IsNullOrEmpty[T any](arr *[]T) bool {
	return arr == nil || len(*arr) == 0
}

func KeySet[T comparable, U any](m *map[T]U) *[]T {
	keys := make([]T, len(*m))
	i := 0
	for k := range *m {
		keys[i] = k
		i++
	}
	return &keys
}

type MapEntry[T comparable, U any] struct {
	Key   T
	Value U
}

func MapEntries[T comparable, U any](m *map[T]U) *[]MapEntry[T, U] {
	entries := make([]MapEntry[T, U], len(*m))
	i := 0
	for k, v := range *m {
		entries[i] = MapEntry[T, U]{k, v}
		i++
	}
	return &entries
}

func MaxBy[T any](arr *[]T, f func(*T) int) T {
	max := (*arr)[0]
	maxValue := f(&max)
	for _, v := range *arr {
		value := f(&v)
		if value > maxValue {
			max = v
			maxValue = value
		}
	}
	return max
}

func JoinToString[T any](arr *[]T, separator string, mapper func(*T) string) string {
	if IsNullOrEmpty(arr) {
		return ""
	}
	mappeed := Map(arr, mapper)

	return strings.Join(mappeed, separator)
}
