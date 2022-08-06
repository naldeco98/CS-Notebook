package main

import (
	"fmt"
	"math/rand"
)

type List []int

func main() {
	rand.Seed(1)
	list := List(rand.Perm(10))

	fmt.Println("Selection Sort")
	fmt.Printf("Before sort: \t%v\n", list)
	list.SelectionSort()
	fmt.Printf("After sort: \t%v\n", list)
}

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

