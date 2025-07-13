//bidirectional communication, main sends jobs , processor reads and process it and and sending it through results channel
package main

import (
	"fmt"
)


// we have two channel as a parameter in the processor, one for receiving(jobs) and other for sending(results)
func processor(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * 2 //sending data through results channel
	}
}

func main() {
	jobs := make(chan int, 5) //created job buffered channel with capacity 5, as a buffered channel lets you send up to 5 values without blocking, unless the buffer is full
	results := make(chan int, 5) //created results channel with capacity 5

	go processor(jobs, results)

	// Send jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	// close(jobs) //it closes the jobs channel, as no more data will be sent and  causing the loop in the processor to stop once all the jobs are received

	// Receive results
	for i := 1; i <= 5; i++ {
		fmt.Println("Result:", <-results)
	}
}
