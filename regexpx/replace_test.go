package regexpx

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestReplaceNString(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		out := ReplaceNString(regexp.MustCompile("bar(.)(x)"), "foobar1xfoobar2x", "baz$1$2", -1)
		assert.Equal(t, out, "foobaz1xfoobaz2x")
	})
	t.Run("none", func(t *testing.T) {
		out := ReplaceNString(regexp.MustCompile("bar(.)(x)"), "foobar1xfoobar2x", "baz$1$2", 0)
		assert.Equal(t, out, "foobar1xfoobar2x")
	})
	t.Run("first", func(t *testing.T) {
		out := ReplaceNString(regexp.MustCompile("bar(.)(x)"), "foobar1xfoobar2x", "baz$1$2", 1)
		assert.Equal(t, out, "foobaz1xfoobar2x")
	})
}
