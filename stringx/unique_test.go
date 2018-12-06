package stringx

import "testing"

func TestUnique(t *testing.T) {
	t.Run("samples", func(t *testing.T) {
		a := []string{"0", "1", "2"}
		ia := Unique(a)
		if len(ia) != 3 {
			t.Fail()
		}
	})
	t.Run("duplicates1", func(t *testing.T) {
		a := []string{"0", "0", "2", "2"}
		ia := Unique(a)
		if len(ia) != 2 {
			t.FailNow()
		}
	})
	t.Run("duplicates2", func(t *testing.T) {
		a := []string{"0", "1", "2", "a", "x", "u", "uu", "uu"}
		ia := Unique(a)
		if len(ia) != 7 {
			t.FailNow()
		}
	})
}

func BenchmarkUnique(b *testing.B) {
	data := make([][]string, b.N)
	for n := 0; n < b.N; n++ {
		data[n] = []string{"0", "1", "2", "a", "x", "u", "uu", "uu", "0", "0", "2", "2", "0", "1", "2"}
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		Unique(data[n])
	}
}