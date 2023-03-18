package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	intMin := min[int64]
	stringMin := min[string]
	testInt := intMin(10, 3)
	testString := stringMin("test", "tes")

	fmt.Println(testInt)
	fmt.Println(testString)

}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
