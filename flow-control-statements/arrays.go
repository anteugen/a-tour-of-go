package main

import "fmt"

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
}