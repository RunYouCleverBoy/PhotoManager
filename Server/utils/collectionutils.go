package utils

func Map[T any, U any](arr *[]T, f func(T) U) []U {
	result := make([]U, len(*arr))
	for i, v := range *arr {
		result[i] = f(v)
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
