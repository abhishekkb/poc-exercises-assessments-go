package goroutines

import (
	"fmt"
	"time"
)

func StoppingGoRoutines() {

	temp := make(chan bool)

	go func() {
		i := 0
		for {
			select {
			case <-temp:
				return
			default:
				fmt.Printf("I'm still running :( , %d \n", i)
				i += 1
			}
		}
	}()

	for i := 0; i < 4; i++ {
		time.Sleep(5 * time.Second)
	}
	temp <- true
}
