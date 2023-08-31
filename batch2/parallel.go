package batch2

import (
	"fmt"
	"sync"
)

type Job struct {
	id    int
	chunk []int
}

func PSum() {
	chunks, numChunks := createChunks(10, createNumSlice())
	sumChan := make(chan int, numChunks)
	defer close(sumChan)
	totalSum := 0
	go func() {
		for sum := range sumChan {
			totalSum += sum
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(numChunks)
	for i := 0; i < numChunks; i++ {
		go sumArray(chunks[i], wg, sumChan)
	}
	wg.Wait()
	fmt.Printf("total sum %d\n", totalSum)

}

func sumArray(a []int, wg *sync.WaitGroup, sumChan chan int) {
	defer wg.Done()
	s := 0
	for _, v := range a {
		s += v
	}
	sumChan <- s
}
func createChunks(chunkSize int, data []int) ([][]int, int) {
	chunks := [][]int{}
	numChunks := len(data) / chunkSize
	for i := 0; i < numChunks; i++ {
		c := i * chunkSize
		chunk := data[c : c+chunkSize]
		chunks = append(chunks, chunk)
	}
	return chunks, numChunks
}

func createNumSlice() []int {
	l := []int{}
	l = append(l, 2)
	l = append(l, 4)
	l = append(l, 21)
	l = append(l, 32)
	l = append(l, 9)

	l = append(l, 12)
	l = append(l, 44)
	l = append(l, 211)
	l = append(l, 320)
	l = append(l, 89)

	l = append(l, 100)
	l = append(l, 74)
	l = append(l, 11)
	l = append(l, 30)
	l = append(l, 63)

	l = append(l, 102)
	l = append(l, 98)
	l = append(l, 740)
	l = append(l, 78)
	l = append(l, 86)

	l = append(l, 54)
	l = append(l, 29)
	l = append(l, 23)
	l = append(l, 65)
	l = append(l, 68)

	l = append(l, 76)
	l = append(l, 87)
	l = append(l, 32)
	l = append(l, 754)
	l = append(l, 567)
	return l
}
