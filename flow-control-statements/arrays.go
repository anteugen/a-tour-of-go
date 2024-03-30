package main

import "fmt"

func createSlice() {
	a := make([]int, 5)
	printSlice("a:", a)

	b := make([]int, 3, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	d = append(d, 1, 2)
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	var a [3]int
	fmt.Println(a)
	a[1] = 3
	fmt.Println(a[1], a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[3], primes)

	var s []int = primes[1:4]
	fmt.Println(s)
	var s2 []int = primes[:]
	s2 = s2[:2]
	fmt.Println(s2, len(s2), cap(s2))

	createSlice()
}