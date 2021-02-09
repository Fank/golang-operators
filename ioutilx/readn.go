package ioutilx

import (
	"io"
	"io/ioutil"
)

// ReadN reads an given amount of bytes and returns it.
func ReadN(r io.Reader, n int64) ([]byte, bool, error) {
	b, err := ioutil.ReadAll(io.LimitReader(r, n+1))
	if err != nil {
		return nil, false, err
	}
	if int64(len(b)) > n {
		return b[:n], true, nil
	}
	return b, false, nil
}
