package interfaces

type CompareFunc[T1 any, T2 any] func(T1, T2) int

func CompareNumbers[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](a, b T) int {
	return int(a - b)
}
