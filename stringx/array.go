package stringx

import "strconv"

func ToIntArray(a []string) ([]int, error) {
	ia := make([]int, len(a))

	el := 0
	for _, v := range a {
		if v == "" {
			continue
		}

		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		ia[el] = i
		el++
	}

	return ia[:el], nil
}
