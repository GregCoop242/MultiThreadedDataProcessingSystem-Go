package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Task struct {
	Name string
}

func worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Worker %d started\n", id)
	for task := range tasks {
		time.Sleep(500 * time.Millisecond) // Simulate processing
		result := fmt.Sprintf("%s processed by worker %d", task.Name, id)
		results <- result
	}
	log.Printf("Worker %d completed\n", id)
}

func main() {
	const numWorkers = 4
	const numTasks = 10

	tasks := make(chan Task, numTasks)
	results := make(chan string, numTasks)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Add tasks
	for i := 0; i < numTasks; i++ {
		tasks <- Task{Name: fmt.Sprintf("Task-%d", i)}
	}
	close(tasks)

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Print results
	fmt.Println("Processed Results:")
	for result := range results {
		fmt.Println(result)
	}
}