package test

import (
	"fmt"
	"sync"
	"testing"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		// Simulate some work
		// You can replace this with your actual task
		result := job * 2
		results <- result
	}
}

func TestConcurent(t *testing.T) {
	numJobs := 10
	numWorkers := 3

	// Create channels for jobs and results
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs to workers
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
