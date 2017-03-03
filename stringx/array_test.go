package stringx

import (
	"testing"
)

func TestToIntArray(t *testing.T) {
	t.Run("samples", func(t *testing.T) {
		a := []string{"0", "1", "2"}
		ToIntArray(a)
	})
	t.Run("empty", func(t *testing.T) {
		a := []string{"0", "", "2"}
		ToIntArray(a)
	})
	t.Run("NaN", func(t *testing.T) {
		a := []string{"0", "a", "2"}
		ToIntArray(a)
	})
}
