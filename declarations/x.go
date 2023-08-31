package declarations

import "fmt"

func VariableDeclarations() {
	var a, b, c = 1, "something", 40.6
	x, y, z := 1, "something", 40.6
	fmt.Printf("a=%d, b=%s, c=%v\n", a, b, c)
	fmt.Printf("x=%d, y=%s, z=%v\n", x, y, z)

}
