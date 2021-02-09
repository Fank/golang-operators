package ioutilx

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadN(t *testing.T) {
	t.Run("9in3out", func(t *testing.T) {
		b, truncate, err := ReadN(bytes.NewReader([]byte("asdasdasd")), 3)
		assert.NoError(t, err)
		assert.Equal(t, truncate, true)
		assert.Equal(t, []byte("asd"), b)
	})
	t.Run("9in9out", func(t *testing.T) {
		b, truncate, err := ReadN(bytes.NewReader([]byte("asdasdasd")), 9)
		assert.NoError(t, err)
		assert.Equal(t, truncate, false)
		assert.Equal(t, []byte("asdasdasd"), b)
	})
	t.Run("3in9out", func(t *testing.T) {
		b, truncate, err := ReadN(bytes.NewReader([]byte("asd")), 9)
		assert.NoError(t, err)
		assert.Equal(t, truncate, false)
		assert.Equal(t, []byte("asd"), b)
	})
	t.Run("0in2out", func(t *testing.T) {
		b, truncate, err := ReadN(bytes.NewReader([]byte("")), 2)
		assert.NoError(t, err)
		assert.Equal(t, truncate, false)
		assert.Equal(t, []byte(nil), b)
	})
	t.Run("nilin2out", func(t *testing.T) {
		b, truncate, err := ReadN(bytes.NewReader(nil), 2)
		assert.NoError(t, err)
		assert.Equal(t, truncate, false)
		assert.Equal(t, []byte(nil), b)
	})
}
