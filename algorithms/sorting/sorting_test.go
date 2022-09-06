package sorting

import "testing"

func BenchmarkSelection(b *testing.B) {

	list := List([]int{1, 6, 2, 34, 6, 7})

	for i := 0; i < b.N; i++ {
		list.SelectionSort()
	}

}
