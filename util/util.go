package util

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Btoa(a bool) string {
	if a {
		return "1"
	}
	return "0"
}
