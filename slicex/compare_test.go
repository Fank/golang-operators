package slicex_test

import (
	"testing"

	"github.com/fank/golang-operators/slicex"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Run("string_1", func(t *testing.T) {
		arr := []string{"A", "b", "0"}
		assert.Equal(t, true, slicex.Contains[string](arr, "A"))
	})
	t.Run("string_2", func(t *testing.T) {
		arr := []string{"A", "b", "0"}
		assert.Equal(t, false, slicex.Contains[string](arr, "ö"))
	})
	t.Run("int64_1", func(t *testing.T) {
		arr := []int64{1, 5, 1000}
		assert.Equal(t, true, slicex.Contains[int64](arr, 1))
	})
	t.Run("int64_2", func(t *testing.T) {
		arr := []int64{1, 5, 1000}
		assert.Equal(t, false, slicex.Contains[int64](arr, 4))
	})
}

func TestUnique(t *testing.T) {
	t.Run("string_1", func(t *testing.T) {
		arr := []string{"A", "b", "0"}
		assert.Equal(t, []string{"A", "b", "0"}, slicex.Unique[string](arr))
	})
	t.Run("string_2", func(t *testing.T) {
		arr := []string{"A", "b", "0", "0"}
		assert.Equal(t, []string{"A", "b", "0"}, slicex.Unique[string](arr))
	})
	t.Run("int64_1", func(t *testing.T) {
		arr := []int64{1, 5, 1000}
		assert.Equal(t, []int64{1, 5, 1000}, slicex.Unique[int64](arr))
	})
	t.Run("int64_2", func(t *testing.T) {
		arr := []int64{1, 5, 1000, 1000}
		assert.Equal(t, []int64{1, 5, 1000}, slicex.Unique[int64](arr))
	})
}
