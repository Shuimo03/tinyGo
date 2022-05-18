package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	array := [3]int{1, 2, 3}
	res := len(array)
	fmt.Println(res)
}
