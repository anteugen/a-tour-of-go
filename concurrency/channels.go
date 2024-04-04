package main

import (
	"fmt"
	"time"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	ch := make(chan int, 6)
	
	for i, _ := range primes {
		ch <- primes[i]
	}

	for i := 0; i < 6; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(<-ch)
	}
}
