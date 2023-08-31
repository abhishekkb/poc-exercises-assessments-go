package batch

import (
	"fmt"
	"sync"
)

func PSumWgChan() {

	l := createNumSlice()
	chunkSize := 5

	chunks := getChunks(chunkSize, l)
	numChunks := len(chunks)

	sumChan := make(chan int, numChunks)
	defer close(sumChan)
	//errChan := make(chan error)
	totalSum := 0
	//start go routine to get values from channel and add it to totalSum
	go func() {
		for v := range sumChan {
			totalSum += v
		}
	}()

	//start workers
	wg := &sync.WaitGroup{}
	wg.Add(numChunks)
	for i := 0; i < numChunks; i++ {
		go chunkWorker(chunks[i], i+1, wg, sumChan)
	}
	wg.Wait()
	fmt.Printf("Total Sum = %d\n", totalSum)
}

func chunkWorker(chunk []int, currentChunkId int, wg *sync.WaitGroup, sumChan chan<- int) {
	defer wg.Done()
	chunkSum := 0
	for _, v := range chunk {
		chunkSum += v
	}
	sumChan <- chunkSum
	fmt.Printf("Chunk Sum = %d, currentChunkId = %d\n", chunkSum, currentChunkId)
}

func getChunks(chunkSize int, l []int) [][]int {
	chunks := [][]int{}
	numChunks := len(l) / chunkSize
	for i := 0; i < numChunks; i++ {
		c := i * chunkSize
		chunk := l[c : c+chunkSize]
		chunks = append(chunks, chunk)
	}
	return chunks
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
