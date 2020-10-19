package main

import (
	"fmt"
)

// Valid Types for Robohashes
var (
	Robot= Type{"robot", 1}
	Monster = Type{"monster", 2}
	NewRobot = Type{"newRobot", 3}
	Cat = Type{"cat", 4}
	Human = Type{"human", 5}
)

type Type struct {
	name string
	setNumber int
}

type Robohash struct {
	name string // can have spaces, but must not be empty
	rType Type
}

func (r Robohash) GetUrl() string {
	// Sprintf : think of it as string-print-format
	return fmt.Sprintf("https://robohash.org/%s.png?set=set%d", r.name, r.rType.setNumber)
}

func (r Robohash) String() string {
	return fmt.Sprintf("[%s] %s", r.rType.name, r.name)
}

func main() {
	r1 := Robohash{"bob", Robot}
	r2 := Robohash{"meow!!", Cat}
	fmt.Println(r1, r1.GetUrl())
	fmt.Println(r2, r2.GetUrl())
	// Problem: people can do things like this
	r3 := Robohash{"how about monkeys", Type{"monkey", 6}}
	Cat.setNumber = 200
	r4 := Robohash{"", Cat}
	fmt.Println(r3, r4, r4.GetUrl())
}
