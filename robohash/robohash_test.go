package robohash

import (
	"fmt"
	"testing"
)

func TestNewRobohash(t *testing.T) {
	_, err := NewRobohash("", Robot)
	if err == nil {
		t.Errorf("NewRobohash function did not throw an error with a blank name")
	}
}

func ExampleNewRobohash() {
	r, err := NewRobohash("a cat named meow", Cat)
	if err == nil {
		fmt.Println(r)
	}
	// Output:
	// [cat] a cat named meow
}
