package intx

import (
	"testing"
)

func TestJoin(t *testing.T) {
	t.Run("number, char and special", func(t *testing.T) {
		a := []int{0, 1, 5, 123456789}
		ia := Join(a, ",")
		if ia != "0,1,5,123456789" {
			t.Fail()
		}
	})
	t.Run("empty", func(t *testing.T) {
		a := []int{}
		ia := Join(a, ",")
		if ia != "" {
			t.Fail()
		}
	})
}
