package stringx

func Unique(stringSlice []string) []string {
	keys := make(map[string]bool)
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