package main

import (
	"fmt"
	"math"
)

type Informations interface {
	Greetings()
	FutureAge()
}

type Person struct {
	firstName string
	age int 
}

func (person Person) Greetings() {
	fmt.Println("Hello,", person.firstName)
}

func (person Person) FutureAge() {
	fmt.Println("In 10 years,", person.firstName, "will be", (person.age + 10))
}

type Animal struct {
	species string
	age     int
}

func (a Animal) Greetings() {
	fmt.Println("Hello, I am a", a.species)
}

func (a Animal) FutureAge() {
	fmt.Println("In 10 years, I will be", a.age+10, "years old")
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case float64:
		roundedUp := math.Ceil(v)
		fmt.Printf("%v rounds up to %v\n", v, int(roundedUp))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var person1 Informations = Person{"Jan", 34}
	var animal1 Informations = Animal{"dog", 5}
	person1.Greetings()
	person1.FutureAge()
	animal1.Greetings()
	animal1.FutureAge()

	do("Hello")
	do(12)
	do(true)
	do(10.6)
}

