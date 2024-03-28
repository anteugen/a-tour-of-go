package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func increment(a *int) {
    *a = *a + 1
}

func main() {
	v1 := Vertex{1, 2}
	fmt.Println(v1)
	fmt.Println(v1.X)
	v1.Y = 4
	fmt.Println(v1)
	p := &v1
	p.X = 1e9
	fmt.Println(v1)

	x := 5
	increment(&x)
	fmt.Println(x)

	v2 := Vertex{X: 1}
	fmt.Println(v2)
	p  = &Vertex{1, 2}
	fmt.Println(p)
}
