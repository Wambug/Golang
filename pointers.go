package main

import "fmt"

func changeValue(str *string) {
	*str = "Changed!"
}

func changeValue1(str string) {
	str = "Changed!"
}

func main() {
	x := 7
	y := &x
	fmt.Println(x, y)
	*y = 8
	fmt.Println(x, y)
	toChange := "hello"
	changeValue(&toChange)
	fmt.Println(toChange)

	var pointer *string = &toChange
	fmt.Println(pointer, &pointer)
}
