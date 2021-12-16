package slicex_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fank/golang-operators/slicex"
)

func TestConvertToAny(t *testing.T) {
	t.Run("int64", func(t *testing.T) {
		a := []int64{0, 1, 5, 123456789}
		ia := slicex.ConvertToAny[int64](a)
		assert.Equal(t, len(ia), 4)
		for _, v := range ia {
			switch v.(type) {
			case int64:
				// OK
			default:
				assert.Fail(t, "wrong type")
			}
		}
	})
	t.Run("string", func(t *testing.T) {
		a := []string{"", "A", "          "}
		ia := slicex.ConvertToAny[string](a)
		assert.Equal(t, len(ia), 3)
		for _, v := range ia {
			switch v.(type) {
			case string:
				// OK
			default:
				assert.Fail(t, "wrong type")
			}
		}
	})
}
