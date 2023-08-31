package channels

import (
	"fmt"
	"time"
)

func access2(ch chan int) {
	time.Sleep(time.Second)
	fmt.Println("start accessing channel")

	for i := range ch {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func BufferedChannels() {
	// only modify this line to define the capacity
	ch := make(chan int, 3)
	defer close(ch)

	go access2(ch)

	for i := 0; i < 9; i++ {
		ch <- i
		fmt.Println("Filled")
	}

	time.Sleep(3 * time.Second)
}
