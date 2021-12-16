package stringx

import (
	"testing"
)

func TestToIntArray(t *testing.T) {
	t.Run("samples", func(t *testing.T) {
		a := []string{"0", "1", "2"}
		ia, err := ToIntArray(a)
		if err != nil {
			t.Fail()
		}
		if len(ia) != 3 {
			t.Fail()
		}
	})
	t.Run("empty", func(t *testing.T) {
		a := []string{"0", "", "2"}
		ia, err := ToIntArray(a)
		if err != nil {
			t.Fail()
		}
		if len(ia) != 2 {
			t.FailNow()
		}
		if ia[0] != 0 {
			t.Fail()
		}
		if ia[1] != 2 {
			t.Fail()
		}
	})
	t.Run("NaN", func(t *testing.T) {
		a := []string{"0", "a", "2"}
		_, err := ToIntArray(a)
		if err == nil {
			t.Fail()
		}
	})
}
