package main

import (
	"algorithms/sorting"
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1)
	list := sorting.List(rand.Perm(10))

	fmt.Println("Selection Sort")
	fmt.Printf("Before sort: \t%v\n", list)
	list.SelectionSort()
	fmt.Printf("After sort: \t%v\n", list)
}
