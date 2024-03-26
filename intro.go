package main

import (
	"fmt"
	"time"
	"math"
)

const i, j int = 1, 2

func main() {
	fmt.Println("Hello, Antoine")
	fmt.Println("The time is", time.Now())
	fmt.Println(add(20, 20))
	a, b := swap("World", "Hello")
	fmt.Println(a, b)
	fmt.Println(split(17))
	c, python, java := true, false, "no!"
	fmt.Println(i + j, c, python, java)
	f := math.Sqrt(float64(i*i + j*j)) 
	var z uint = uint(f)
	fmt.Println(f, z)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
