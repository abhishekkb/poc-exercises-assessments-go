package batch

import (
	"fmt"
)

type Job struct {
	id    int
	chunk []int
}

func PSumW() {
	chunks := getChunks(5, createNumSlice())
	fmt.Printf("chunks - %v\n", chunks)

	jobs := []Job{}
	for i, chunk := range chunks {
		jobs = append(jobs, Job{i, chunk})
	}

	numChunks := len(chunks)
	jobsChan := make(chan Job, numChunks)
	resultsChan := make(chan int, numChunks)

	totalSum := 0
	numWorkers := 3
	// start fixed number of workers, numWorkers < number of chunks or jobs
	for w := 0; w < numWorkers; w++ {
		go worker(w, jobsChan, resultsChan)
	}

	//add jobs to jobs channel
	for j := 0; j < numChunks; j++ {
		jobsChan <- jobs[j]
	}
	close(jobsChan)

	for r := 0; r < numChunks; r++ {
		totalSum += <-resultsChan
	}
	close(resultsChan)
	fmt.Printf("Total sum %d \n", totalSum)

}

// worker(workerId, jobs, results), this loop on jobs channel
func worker(id int, jobs <-chan Job, results chan<- int) {
	for job := range jobs {
		s := sum(job.chunk)
		fmt.Printf("workerId %d, jobId %d, chunk %v, sum of chunk %d\n", id, job.id, job.chunk, s)
		results <- s
	}
}

func sum(chunk []int) int {
	//fmt.Printf("sum :: chunk %v \n", chunk)
	s := 0
	for i := 0; i < len(chunk); i++ {
		s += chunk[i]
	}
	return s
}
