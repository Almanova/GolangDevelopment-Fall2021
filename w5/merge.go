package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func generator(n int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- n
			time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		}

		close(ch)
	}()

	return ch
}

func merge(cs ...<-chan int) <-chan int {
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	for _, c := range cs {
		wg.Add(1)

		go func(localC <-chan int) {
			defer wg.Done()

			for in := range localC {
				ch <- in
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	rand.Seed(time.Now().UnixNano())

	merged := merge(generator(2), generator(3))

	for v := range merged {
		fmt.Println(v)
	}
}