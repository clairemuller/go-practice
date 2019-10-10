package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	age  int
}

func (p Person) yearBorn(age int) int {
	year := time.Now().Year()
	birthYear := year - age
	return birthYear
}

func main() {
	p := Person{"Claire", 28}
	result := p.yearBorn(p.age)
	fmt.Println(result)
}
