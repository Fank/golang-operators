package intx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExcept(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		diff := Except([]int{0, 1, 2}, []int{0, 1, 2})
		if len(diff) != 0 {
			t.Fail()
		}
	})
	t.Run("test2", func(t *testing.T) {
		diff := Except([]int{1, 2}, []int{0, 2})
		assert.Equal(t, []int{1}, diff)
	})
	t.Run("duplicates2", func(t *testing.T) {
		diff := Except([]int{0, 1, 2}, []int{0, 2})
		assert.Equal(t, []int{1}, diff)
	})
}
