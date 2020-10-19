package main

import (
	"fmt"
)

type Robohash struct {
	name string // can have spaces, but must not be empty
	rType int // 1 - robot, 2 - monster, 3- new robot, 4 - cat, 5 - human
}

func (r Robohash) GetUrl() string {
	// Sprintf : think of it as string-print-format
	return fmt.Sprintf("https://robohash.org/%s.png?set=set%d", r.name,
		r.rType)
}

func (r Robohash) String() string {
	var rTypeName string
	switch (r.rType) {
	case 1:
		rTypeName = "robot"
	case 2:
		rTypeName = "monster"
	case 3:
		rTypeName = "newRobot"
	case 4:
		rTypeName = "cat"
	case 5:
		rTypeName = "human"
	}
	return fmt.Sprintf("[%s] %s", rTypeName, r.name)
}

func main() {
	r1 := Robohash{"bob", 1} // robot
	r2 := Robohash{"meow!!", 4} // cat
	fmt.Println(r1, r1.GetUrl())
	fmt.Println(r2, r2.GetUrl())
}
