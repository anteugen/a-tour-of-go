package main

import (
	"fmt"
	"time"
	"sync"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func count(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go say("world", &wg)

	wg.Add(1)
	go count(&wg)

	wg.Add(1)
	go say("hello", &wg)

	wg.Wait()
}
