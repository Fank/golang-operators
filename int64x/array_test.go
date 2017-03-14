package int64x

import (
	"testing"
)

func TestToInterfaceArray(t *testing.T) {
	t.Run("number, char and special", func(t *testing.T) {
		a := []int64{0, 1, 5, 123456789}
		ia := ToInterfaceArray(a)
		if len(ia) != 4 {
			t.Fail()
		}
		for _, v := range ia {
			switch v.(type) {
			case int64:
				// OK
			default:
				t.Fail()
			}
		}
	})
}
