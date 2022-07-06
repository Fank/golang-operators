package slicex

import "golang.org/x/exp/constraints"

// Except return elements in a which are not present in b
func Except[T constraints.Ordered](a, b []T) []T {
	var diff []T

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
