package robohash

import (
	"fmt"
)

// Valid Types for Robohashes
var (
	Robot = Type{"robot", 1}
	Monster = Type{"monster", 2}
	NewRobot = Type{"newRobot", 3}
	Cat = Type{"cat", 4}
	Human = Type{"human", 5}
)

type Type struct {
	name string
	setNumber int
}

func (t Type) Name() string {
	return t.name
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

func NewRobohash(name string, rType Type) (Robohash, error){
	if name == "" {
		return Robohash{}, fmt.Errorf("Name cannot be an empty string!")
	}
	return Robohash{name, rType}, nil
}
