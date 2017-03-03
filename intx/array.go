package intx

// ToInterfaceArray converter
func ToInterfaceArray(a []int) []interface{} {
	interfaceArray := make([]interface{}, len(a))

	for i, v := range a {
		interfaceArray[i] = v
	}

	return interfaceArray
}
