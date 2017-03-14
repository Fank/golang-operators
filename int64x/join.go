package int64x

import (
	"strconv"
	"strings"
)

// Join int array with seperator
func Join(a []int64, sep string) string {
	switch len(a) {
	case 0:
		return ""
	}

	sa := make([]string, len(a))
	for i, v := range a {
		sa[i] = strconv.FormatInt(v, 10)
	}

	return strings.Join(sa, sep)
}
