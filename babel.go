package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

func r(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
func main() {
	a := r(1, 300)
	fmt.Println("This will be a book of babel simulation")
	fmt.Println("", a)
}
