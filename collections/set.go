package collections

import "fmt"

func CreateSetOfString() {
	s := map[string]struct{}{}
	s["a"] = struct{}{}
	s["b"] = struct{}{}
	s["c"] = struct{}{}
	s["d"] = struct{}{}
	s["d"] = struct{}{}
	fmt.Printf("map = %v\n", s)
}
