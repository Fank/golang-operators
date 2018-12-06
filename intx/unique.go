package intx

func Unique(stringSlice []int) []int {
	keys := make(map[int]bool)
	lastIndex := 0
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			stringSlice[lastIndex] = entry
			lastIndex++
		}
	}
	return stringSlice[:lastIndex]
}
