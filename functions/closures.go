package functions

import "fmt"

func getAdder() func(int) int {
	x := 0 // this will be referenced outside of this function

	return func(i int) int {
		x += i
		return x
	}
}

func RunClosures() {
	add, sub := getAdder(), getAdder()
	c := 10
	for i := 0; i < 100; i++ {
		fmt.Printf("values for add() and sub() are %d, %d \n", add(i+c), sub(-1*(i+c)))
	}

	fmt.Println("\n\n\n==================Running another time by assigning new func refs to the vars==================")
	add, sub = getAdder(), getAdder()

	for i := 0; i < 100; i++ {
		fmt.Printf("values for add() and sub() are %d, %d \n", add(i+c), sub(-1*(i+c)))
	}
}
