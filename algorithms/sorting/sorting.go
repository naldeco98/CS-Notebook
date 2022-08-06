package sorting

type List []int

func (elements *List) SelectionSort() {
	min := 0
	length := len(*elements)
	for i := 0; i < length; i++ {
		min = i
		for j := i; j < length; j++ {
			if (*elements)[j] < (*elements)[min] {
				min = j
			}
		}
		(*elements)[i], (*elements)[min] = (*elements)[min], (*elements)[i]
	}
}
