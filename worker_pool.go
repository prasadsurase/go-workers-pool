package main

import (
	"fmt"
	"sync"
)

// WorkerPool holds the configuration to manage the workers.
type WorkerPool struct {
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// Worker() method retrieves the tasks from the channel and processes them.
func (wp *WorkerPool) Worker() {
	fmt.Println("--------------------------------------------")
	fmt.Println("Worker running. Checking channel...")
	for task := range wp.tasksChan {
		fmt.Println("Retrieved task. Processing...")
		task.Process()
		wp.wg.Done()
	}
}

// Run methods starts N workers as per the worker pool's configuration.
func (wp *WorkerPool) Run() {
	for i := 0; i < wp.concurrency; i++ {
		fmt.Printf("Starting worker %d\n", i+1)
		go wp.Worker()
	}
}
