package main

import (
	"fmt"
)

func _main() {
	arr := [3]int{4, 5, 6}
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)
}
