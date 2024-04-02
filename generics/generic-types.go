package main

import "fmt"

type List[T any] struct {
    next *List[T]
    val  T
}

func (l *List[T]) Add(value T) {
    newNode := &List[T]{val: value, next: nil}
    if l == nil {
        l = newNode
    } else {
        current := l
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}


func main() {
    head := &List[int]{val: 0}
   
    for i := 1; i <= 5; i++ {
        head.Add(i * 2)
    }

	head2 := &List[string]{val: "Hello"}
   
	stringList := []string{"Reader,", "how", "are", "you?"}
	
    for i := 0; i <= 3; i++ {
        head2.Add(stringList[i])
    }

    for node := head; node != nil; node = node.next {
        fmt.Println(node.val)
    }

	for node := head2; node != nil; node = node.next {
        fmt.Println(node.val)
    }
}
