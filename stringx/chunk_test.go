package stringx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunkHuman(t *testing.T) {
	t.Run("nolimit singleline", func(t *testing.T) {
		chunk, truncated := ChunkHuman(
			`Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam`,
			-1,
			-1,
			)
		assert.Equal(t, truncated, false)
		assert.Equal(t, chunk, []string{
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		})
	})
	t.Run("nolimit multiline", func(t *testing.T) {
		chunk, truncated := ChunkHuman(
			`Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam`,
			-1,
			-1,
		)
		assert.Equal(t, truncated, false)
		assert.Equal(t, chunk, []string{
			`Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam`,
		})
	})
	t.Run("fit singleline", func(t *testing.T) {
		chunk, truncated := ChunkHuman(
			`Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam`,
			1,
			130,
		)
		assert.Equal(t, truncated, false)
		assert.Equal(t, chunk, []string{
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		})
	})
	t.Run("fit multiline", func(t *testing.T) {
		chunk, truncated := ChunkHuman(
			`Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam`,
			3,
			130,
		)
		assert.Equal(t, truncated, false)
		assert.Equal(t, chunk, []string{
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		})
	})
	t.Run("overflow multiline", func(t *testing.T) {
		chunk, truncated := ChunkHuman(
			`Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyamaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyamaaaaaaaaaaaaaaaaaaaaaaaaaaa`,
			2,
			130,
		)
		assert.Equal(t, truncated, true)
		assert.Equal(t, chunk, []string{
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut ",
		})
	})
	t.Run("10x10 fit", func(t *testing.T) {
		chunk, truncated := ChunkHuman(
			`Lorem ipsum d
Lorem ipsum dolor sit`,
			10,
			10,
		)
		assert.Equal(t, truncated, false)
		assert.Equal(t, chunk, []string{
			"Lorem",
			"ipsum d",
			"Lorem",
			"ipsum",
			"dolor sit",
		})
	})
	t.Run("nolimit singleline", func(t *testing.T) {
		chunk, truncated := ChunkHuman(
			`Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam`,
			3,
			50,
		)
		assert.Equal(t, truncated, false)
		assert.Equal(t, chunk, []string{
			"Lorem ipsum dolor sit amet, consetetur sadipscing",
			"elitr, sed diam nonumy eirmod tempor invidunt ut",
			"labore et dolore magna aliquyam",
		})
	})
	t.Run("nolimit singleline", func(t *testing.T) {
		chunk, truncated := ChunkHuman(
			`Bestellanforderungsnummer 1152149880
 LR-Nr. 6102528
 Kostenstelle 20000520`,
			4,
			35,
		)
		assert.Equal(t, truncated, false)
		assert.Equal(t, chunk, []string{
			"Lorem ipsum dolor sit amet, consetetur sadipscing",
			"elitr, sed diam nonumy eirmod tempor invidunt ut",
			"labore et dolore magna aliquyam",
		})
	})
	t.Run("suffix spaces", func(t *testing.T) {
		chunk, truncated := ChunkHuman(
			`Bestellung Nr. 12345
Kommission: Hr. Foo
`,
			4,
			35,
		)
		assert.Equal(t, truncated, false)
		assert.Equal(t, chunk, []string{
			"Bestellung Nr. 12345",
			"Kommission: Hr. Foo",
		})
	})
}
