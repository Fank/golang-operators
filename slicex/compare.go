package slicex

// Contains checks if the val is found in arr
func Contains[T comparable](arr []T, val T) bool {
	found := false

	for _, v := range arr {
		if v == val {
			found = true
			break
		}
	}

	return found
}

// Unique removes duplicates from arr
func Unique[T comparable](arr []T) []T {
	keys := make(map[T]bool)
	lastIndex := 0

	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			arr[lastIndex] = entry
			lastIndex++
		}
	}

	return arr[:lastIndex]
}
