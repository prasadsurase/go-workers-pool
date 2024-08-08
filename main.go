package main

import (
	"fmt"
	"runtime"
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
		// concurrency: 3,
		concurrency: runtime.NumCPU(),
		tasksChan:   make(chan Task, 1), // buffered channel size should never be more than 1.
		doneChan:    make(chan bool),
	}

	go func() {
		fmt.Printf("Starting workers...\n")
		wp.Run()
	}()

	go func() {
		fmt.Printf("Loading tasks...\n")
		// Load tasks in the channel.
		// This can be reading a large file or reading data from a queue
		for i := 0; i < len(tasks); i++ {
			fmt.Printf("Loading task %d in tasks channel\n", i+1)
			wp.tasksChan <- tasks[i]
		}
		close(wp.tasksChan)
		wp.doneChan <- true
	}()

	// wait until done channel has some data.
	<-wp.doneChan

	fmt.Println("Main finished as no tasks")
}
