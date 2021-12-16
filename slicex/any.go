package slicex

// ConvertToAny create an any array
func ConvertToAny[T any](a []T) []any {
	interfaceArray := make([]any, len(a))

	for i, v := range a {
		interfaceArray[i] = v
	}

	return interfaceArray
}
