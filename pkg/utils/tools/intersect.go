package tools

func Contains[T comparable](collection []T, element T) bool {
	for _, t := range collection {
		if t == element {
			return true
		}
	}
	return false
}

// Empty returns an empty value.
func Empty[T any]() T {
	var t T
	return t
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
