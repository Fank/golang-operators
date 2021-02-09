package ioutilx

import (
	"io"
)

// ReadN reads an given amount of bytes and returns it.
func ReadN(r io.Reader, n int) ([]byte, bool, error) {
	buf := make([]byte, n+1)
	v, err := r.Read(buf)
	if err == io.EOF {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	index := v
	if v > n {
		index = n
	}
	return buf[:index], v > n, nil
}
