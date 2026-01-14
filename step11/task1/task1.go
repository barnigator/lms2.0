package main

type MyConst interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Sum[T MyConst](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}
