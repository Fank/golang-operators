package intx

func Except(a []int, b []int) []int {
	var diff []int

	exists := false
	for _, v := range a {
		exists = false
		for _, vv := range b {
			exists = v == vv
			if exists {
				break
			}
		}
		if !exists {
			diff = append(diff, v)
		}
	}

	return diff
}
