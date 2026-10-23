package main

import "fmt"

func main() {
	values := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 19,
	}

	v, b := Primeiro[int](values)

	fmt.Printf("%v, %v", v, b)

	mapStruct := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Println(mapStruct["a"])

	out := Unicos(values)
	fmt.Println(out)
}
