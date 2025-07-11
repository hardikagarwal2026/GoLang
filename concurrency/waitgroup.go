// to wait for multiple goroutines to finish, we use goroutines
// A WaitGroup in Go is a concurrency synchronization primitive provided by the sync package.
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) { // waitgroup is passsed as pointer in the function parameter
	defer wg.Done() //// Mark this worker as done when it exits, we are using  defer wg.Done() at the start of the goroutine to ensure it runs even if the function panics or returns early.
	fmt.Printf("Hello we are starting, %d\n", id)
	time.Sleep(time.Second) //do to some work
	fmt.Printf("we are done, %d\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)         //// increment counter by 1, That means I am starting one more goroutine.
		go worker(i, &wg) //waitgroup is being passed as a pointer
	}

	wg.Wait() //Wait for all 3 workers to finished
	fmt.Println("All workers finished!")
}
