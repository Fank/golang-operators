package stringx

import (
	"strings"
)

func ChunkHuman(str string, rows, chars int) ([]string, bool) {
	if chars < 0 && rows < 0 {
		return []string{str}, false
	}

	lines := strings.Split(str, "\n")
	charsOverflow := false
	charCount := 0
	i := 0
	for _, v := range lines {
		if v == "" {
			continue
		}

		charCount += len(v)
		if len(v) > chars {
			charsOverflow = true
		}

		lines[i] = v
		i++
	}
	lines = lines[:i]

	if len(lines) <= rows && !charsOverflow {
		return lines, false
	}

	noNewLineString := strings.Replace(strings.Replace(str, "\r", "", -1), "\n", " ", -1)
	var chunks []string
	if charCount > rows*chars {
		for i := 0; i < len(noNewLineString); i += chars {
			nn := i + chars
			if nn > len(noNewLineString) {
				nn = len(noNewLineString)
			}
			chunks = append(chunks, string(noNewLineString[i:nn]))
			if len(chunks) == rows {
				break
			}
		}
		return chunks, true
	}

	buffer := ""
	for i, v := range strings.Split(noNewLineString, " ") {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}

		if len(buffer) + 1 + len(v) > chars {
			chunks = append(chunks, buffer)
			if len(chunks) == rows {
				break
			}
			buffer = ""
		} else if i > 0 {
			buffer += " "
		}
		buffer += v
	}
	if buffer != "" {
		chunks = append(chunks, buffer)
	}

	return chunks, false
}
