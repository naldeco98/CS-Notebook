// Exercie 1 - Echo function
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	// var s, sep string
	// for _, arg := range os.Args[0:] {
	// 	s += sep + arg
	// }
	// fmt.Println(s)
	// fmt.Println("took:", time.Now().Sub(now))
	// now = time.Now()
	fmt.Println(strings.Join(os.Args[0:], ""))
	fmt.Println("took:", time.Now().Sub(now))
}
