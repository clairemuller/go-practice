package main

import (
	"fmt"
	"time"
)

// https://golangbot.com/structs/
// creates named struct
// versus anonymous struct, which would be assigned to a variable
type Person struct {
	name string
	age  int
}

// https://golangbot.com/methods/
// creates a method, which is just a function
// with a receiver argument btwn the func keyword and the method name
func (p Person) yearBorn() int {
	currYear := time.Now().Year()
	birthYear := currYear - p.age
	return birthYear
}

func main() {
	p1 := Person{"Claire", 28}

	fmt.Printf("%s is %d and was born in %d\n", p1.name, p1.age, p1.yearBorn())
	
	p2 := Person{
		name: "Kristen",
		age:  36,
	}

	fmt.Printf("%s is %d and was born in %d\n", p2.name, p2.age, p2.yearBorn())
}
