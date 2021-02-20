package main

import (
	"fmt"
)

func main() {
	var a []int = []int{1, 3, 4, 56, 7, 12, 4, 6}

	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}
}
