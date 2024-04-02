package main

import "fmt"

type Number interface {
	int | float64 | string
}

func Adder[T Number](a, b T) T {
	return a + b
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	intList := []int{10, 15, 20, 25}
	fmt.Println(Index(intList, 15))
	fmt.Println("Sum:",Adder(25.45, 4))
	fmt.Println("Full sentence:", Adder("Hello, ", "Reader"))
}
