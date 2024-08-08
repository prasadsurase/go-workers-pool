package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())

	tasks := []Task{
		&EmailProcessor{email: "prasadsurase+1@gmail.com"},
		&ImageProcessor{imageUrl: "/images/sample1.jpeg"},
		&EmailProcessor{email: "prasadsurase+2@gmail.com"},
		&ImageProcessor{imageUrl: "/images/sample2.jpeg"},
		&EmailProcessor{email: "prasadsurase+3@gmail.com"},
		&ImageProcessor{imageUrl: "/images/sample3.jpeg"},
		&ImageProcessor{imageUrl: "/images/sample4.jpeg"},
		&EmailProcessor{email: "prasadsurase+4@gmail.com"},
		&ImageProcessor{imageUrl: "/images/sample5.jpeg"},
		&EmailProcessor{email: "prasadsurase+5@gmail.com"},
		&ImageProcessor{imageUrl: "/images/sample6.jpeg"},
	}

	wp := WorkerPool{
		concurrency: 3,
		// concurrency: runtime.NumCPU(),
		tasksChan: make(chan Task, 1), // buffered channel size should never be more than 1.
		wg:        sync.WaitGroup{},
	}

	go func() {
		fmt.Printf("Loading tasks...\n")
		// Load tasks in the channel.
		// This can be reading a large file or reading data from a queue
		for i := 0; i < len(tasks); i++ {
			fmt.Printf("Loading task %d in tasks channel\n", i+1)
			wp.wg.Add(1)
			wp.tasksChan <- tasks[i]
		}
		close(wp.tasksChan)
	}()

	go func() {
		fmt.Printf("Starting workers...\n")
		wp.Run()
	}()

	time.Sleep(5 * time.Millisecond) //without this statement, program would exit as go routines cannot populate the tasks
	wp.wg.Wait()
	fmt.Println("Main finished as no tasks")
}
