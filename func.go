package algo

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinStr(a, b string) string {
	if a < b {
		return a
	}
	return b
}

func MaxStr(a, b string) string {
	if a > b {
		return a
	}
	return b
}
