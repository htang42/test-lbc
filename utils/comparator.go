package utils

func Max[T int](a, b T) T {
	if a < b {
		return b
	}
	return a
}
