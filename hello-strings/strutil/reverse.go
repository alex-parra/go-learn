package strutil

// Reverse returns the provided string and return it reversed
func Reverse(s string) string {
	letters := []rune(s)

	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		letters[i], letters[j] = letters[j], letters[i]
	}

	return string(letters)
}
