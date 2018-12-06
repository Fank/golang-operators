package intx

import "testing"

func TestUnique(t *testing.T) {
	t.Run("samples", func(t *testing.T) {
		a := []int{0, 1, 2}
		ia := Unique(a)
		if len(ia) != 3 {
			t.Fail()
		}
	})
	t.Run("duplicates1", func(t *testing.T) {
		a := []int{0, 0, 2, 2}
		ia := Unique(a)
		if len(ia) != 2 {
			t.FailNow()
		}
	})
	t.Run("duplicates2", func(t *testing.T) {
		a := []int{0, 1 ,2 ,3 ,4 ,4 ,5 ,6 ,0, 100, 21323, 151, 10, 0}
		ia := Unique(a)
		if len(ia) != 7 {
			t.FailNow()
		}
	})
}

func BenchmarkUnique(b *testing.B) {
	data := make([][]int, b.N)
	for n := 0; n < b.N; n++ {
		data[n] = []int{0, 1, 2,0, 0, 2, 2, 0, 1 ,2 ,3 ,4 ,4 ,5 ,6 ,0, 100, 21323, 151, 10, 0}
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		Unique(data[n])
	}
}