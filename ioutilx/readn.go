package ioutilx

import (
	"io"
)

// ReadN reads an given amount of bytes and returns it.
func ReadN(r io.Reader, n int) ([]byte, bool, error) {
	var s []byte
	size := 0

	buf := make([]byte, 1024)
	for {
		v, err := r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, false, err
		}

		s = append(s, buf...)
		size += v
		if size >= n {
			break
		}
	}

	index := size
	if size > n {
		index = n
	}

	return s[:index], size > n, nil
}
