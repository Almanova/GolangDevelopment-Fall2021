package main

import (
    "fmt"
    "sync"
    "time"
	"math/rand"
)

var wg sync.WaitGroup

func worker(id int, tasks <-chan int, results chan<- int) {
	for i := range tasks {
		fmt.Printf("Worker %d started executing task %d\n", id, i)

		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)

		fmt.Printf("Worker %d finished with the task %d\n", id, i)
		results <- i
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	taskCnt := 6
	workerCnt := 2
    tasks := make(chan int, taskCnt)
    results := make(chan int, taskCnt)

    for i := 1; i <= workerCnt; i++ {
        wg.Add(1)

        go func(localI int) {
            defer wg.Done()
            worker(localI, tasks, results)
        }(i)
    }

	for i := 1; i <= taskCnt; i++ {
        tasks <- i
    }
    close(tasks)

    for i := 1; i <= taskCnt; i++ {
        fmt.Printf("Task %d was executed\n", <-results)
    }

	wg.Wait()
}