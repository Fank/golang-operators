package stringx

// Limit limits the string to the provided length
func Limit(s string, length int) string {
	if len(s) <= length {
		return s
	}

	return s[:length]
}
