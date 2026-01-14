package main

func Filter[T any](arr []T, predicate func(T) bool) []T {
	res := make([]T, 0, len(arr))
	for _, val := range arr {
		if predicate(val) {
			res = append(res, val)
		}
	}
	return res
}
