package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	fmt.Println(mp["apple"])
	//delete(mp, "tim")
	fmt.Println(mp)

	val, ok := mp["tim"]
	fmt.Println(val, ok)
}
