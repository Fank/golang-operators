package stringx

import "testing"

func TestSliceAppendIfMissing(t *testing.T) {
	t.Run("samples", func(t *testing.T) {
		a := []string{"0", "1", "2", "2", "3"}
		b := []string{}

		for _, v := range a {
			b = SliceAppendIfMissing(b, v)
		}

		if len(b) != 4 {
			t.Fail()
		}
	})
}