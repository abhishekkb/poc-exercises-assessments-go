package ducktyping

import "fmt"

//duck typing or structural typing

type Animal interface {
	eat() bool
	walk(x int, y int) (int, int)
	getx() int
	gety() int
}

type Dog struct {
	posx int
	posy int
}

func (d *Dog) eat() bool {
	return true
}

func (d *Dog) walk(x, y int) (int, int) {
	d.posx += x
	d.posy += y
	return d.posx, d.posy
}

func (d *Dog) getx() int {
	return d.posx
}

func (d *Dog) gety() int {
	return d.posy
}

func WalkTheDog() {
	var a Animal = &Dog{0, 0}

	for i := 0; i < 100; i++ {
		fmt.Printf("current pos %d, %d\n", a.getx(), a.gety())
		a.walk(i, i+10)
		fmt.Printf("current pos after walking %d, %d\n", a.getx(), a.gety())
	}
}
