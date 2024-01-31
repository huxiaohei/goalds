package comparator

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Signed | Unsigned
}

type Ordered interface {
	Integer | Float | ~string
}

// Comparator is a function that compares two elements.
// It returns a negative integer if a < b, zero if a == b, or a positive integer if a > b.
type Comparator[T any] func(a, b T) int

func OrderedTypeCmp[T Ordered](a, b T) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func Reverse[T any](cmp Comparator[T]) Comparator[T] {
	return func(a, b T) int {
		return cmp(b, a)
	}
}
