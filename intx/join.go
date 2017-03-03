package intx

import (
	"strconv"
	"strings"
)

// Join int array with seperator
func Join(a []int, sep string) string {
	switch len(a) {
	case 0:
		return ""
	}

	sa := make([]string, len(a))
	for i, v := range a {
		sa[i] = strconv.Itoa(v)
	}

	return strings.Join(sa, sep)
}
