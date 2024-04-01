package main

import (
	"fmt"
)

func sayGreetings(n string) {
	fmt.Println("Good morning,", n)
}

func cycleNames(names []string, f func(string)) {

	for _, v := range names {
		f(v)
	}
}

func adder() func(int) int {
	sum := 5
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sayGreetings("hank")
	cycleNames([]string{"François", "Philippe", "Martin"}, sayGreetings)

	add := adder()
	fmt.Println(add(5))
}