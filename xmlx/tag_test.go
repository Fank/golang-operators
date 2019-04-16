package xmlx

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRootTag(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		tag, err := RootTag(bytes.NewReader([]byte("<foo></foo>")))
		assert.NoError(t, err, "")
		assert.Equal(t, "foo", tag)
	})
	t.Run("with attribute", func(t *testing.T) {
		tag, err := RootTag(bytes.NewReader([]byte(`<foo type="asd"></foo>`)))
		assert.NoError(t, err, "")
		assert.Equal(t, "foo", tag)
	})
	t.Run("encoding", func(t *testing.T) {
		tag, err := RootTag(bytes.NewReader([]byte(`<?xml version="1.0" encoding="ISO-8859-1"?><foo></foo>`)))
		assert.NoError(t, err, "")
		assert.Equal(t, "foo", tag)
	})

}
