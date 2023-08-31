package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func UsingSignals() {

	sig := make(chan struct{})
	defer close(sig)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go receiveSignal(sig, wg)
	time.Sleep(2 * time.Second)
	sig <- struct{}{}
	wg.Wait()
	fmt.Printf("closing, time = %v\n", time.Now())
}

func receiveSignal(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("in receiveSignal func, time = %v \n", time.Now())
	<-ch
	fmt.Printf("received signal, time = %v\n", time.Now())
}
