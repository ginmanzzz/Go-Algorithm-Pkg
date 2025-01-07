package algo

type ComparableType interface {
	~int | ~float64 | ~string | ~int32 | ~float32
}

func Max [T ComparableType] (a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min [T ComparableType] (a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Swap [T any] (a, b *T) {
	*a, *b = *b, *a
}

func IntAbs (a int) int {
    if a >= 0 {
        return a
    }
    return -a
}
