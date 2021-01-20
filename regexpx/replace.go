package regexpx

import "regexp"

func ReplaceNString(regex *regexp.Regexp, src, repl string, n int) string {
	i := 0
	return regex.ReplaceAllStringFunc(src, func(a string) string {
		if n >= 0 && i >= n {
			return a
		}
		i++
		return regex.ReplaceAllString(a, repl)
	})
}
