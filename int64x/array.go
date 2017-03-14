package int64x

// ToInterfaceArray converter
func ToInterfaceArray(a []int64) []interface{} {
	interfaceArray := make([]interface{}, len(a))

	for i, v := range a {
		interfaceArray[i] = v
	}

	return interfaceArray
}
