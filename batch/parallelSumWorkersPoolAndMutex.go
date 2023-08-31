package batch

import (
	"fmt"
	"sync"
)

func PSumWMutex() {
	chunks := getChunks(5, createNumSlice())
	fmt.Printf("chunks - %v\n", chunks)

	mu := &sync.Mutex{}

	jobs := []Job{}
	for i, chunk := range chunks {
		jobs = append(jobs, Job{i, chunk})
	}

	numChunks := len(chunks)
	jobsChan := make(chan Job, numChunks)
	//resultsChan := make(chan int, numChunks)

	zilch := 0
	totalSum := &zilch
	numWorkers := 3
	//add jobs to jobs channel
	for j := 0; j < numChunks; j++ {
		jobsChan <- jobs[j]
	}
	// start fixed number of workers, numWorkers < number of chunks or jobs
	wg := &sync.WaitGroup{}
	for w := 0; w < numWorkers; w++ {
		go worker2(w, jobsChan, mu, totalSum, wg)
	}
	wg.Wait()
	close(jobsChan)

	fmt.Printf("Total sum %d \n", *totalSum)

}

// worker(workerId, jobs, results), this loop on jobs channel
func worker2(id int, jobs <-chan Job, mu *sync.Mutex, ts *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		s := sum(job.chunk)
		fmt.Printf("workerId %d, jobId %d, chunk %v, sum of chunk %d\n", id, job.id, job.chunk, s)
		mu.Lock()
		fmt.Printf("workerId %d, jobId %d, mutex locked\n", id, job.id)
		temp := *ts + s
		ts = &temp
		mu.Unlock()
		fmt.Printf("workerId %d, jobId %d, mutex unlocked\n", id, job.id)
	}
}
