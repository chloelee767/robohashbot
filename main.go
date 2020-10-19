package main

import (
	"fmt"
	"github.com/chloelee767/robohashbot/robohash"
)


func main() {
	r1, err1:= robohash.NewRobohash("bob", robohash.Robot)
	r2, err2 := robohash.NewRobohash("meow!!", robohash.Cat)
	r4, err4 := robohash.NewRobohash("", robohash.Cat)

	// we can now check whether we have created a valid Robohash
	if err1 == nil {
		fmt.Println(r1, r1.GetUrl())
	}

	if err2 == nil {
		fmt.Println(r2, r2.GetUrl())
	}

	if err4 == nil {
		fmt.Println(r4, r4.GetUrl())
	} else {
		fmt.Println("err4:", err4)
	}

	// All of this will not compile now as the struct fields are unexported:
	// r1 := robohash.Robohash{"bob", robohash.Robot}
	// r2 := robohash.Robohash{"meow!!", robohash.Cat}
	// r3 := robohash.Robohash{"how about monkeys", roboash.Type{"monkey", 6}}
	// robohash.Cat.setNumber = 200

}
