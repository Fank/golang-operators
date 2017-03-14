package int64x

import (
	"testing"
)

func TestJoin(t *testing.T) {
	t.Run("number, char and special", func(t *testing.T) {
		a := []int64{0, 1, 5, 123456789}
		ia := Join(a, ",")
		if ia != "0,1,5,123456789" {
			t.Fail()
		}
	})
	t.Run("empty", func(t *testing.T) {
		a := []int64{}
		ia := Join(a, ",")
		if ia != "" {
			t.Fail()
		}
	})
}
