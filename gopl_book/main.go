package main

import "fmt"

func main() {
	newMap := make(map[string]string)
	fmt.Printf("%v\n", newMap)
	myFunc(newMap)
	fmt.Printf("%v\n", newMap)
}

func myFunc(m map[string]string) {
	m["Hello"] = "World"
	fmt.Printf("%v\n", m)
}
