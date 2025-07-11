// package main

// import "fmt"
// // import "time"

// func A(channel1 <-chan int, channel2 chan<- int){
//      for val := range channel1 {
// 		channel2 <- val * 2
// 	 }
// 	 close(channel2)

// }

// func B(channel2 <-chan int, channel3 chan<- int){
//      for val := range channel2 {
// 		// time.Sleep(time.Second)	
// 		channel3 <- val * 2
// 	 }
// 	 close(channel3)
// }


// func main(){
// 	noofjobs := 5
// 	channel1 := make(chan int, noofjobs)
// 	channel2 := make(chan int, noofjobs)
// 	channel3 := make(chan int, noofjobs)


// 	//initializing the 1-1 worker
// 	// go A(channel1, channel2)
// 	// go B(channel2, channel3)


// 	//initializing 3-3 workers
// 	for i:=0; i< 3 ; i++ {
// 		go A(channel1, channel2)
// 	}

// 	for i:=0; i< 3 ; i++ {
// 		go B(channel2, channel3)
// 	}

// 	//sending jobs to channel1
// 	for i := 1; i<= noofjobs;i++ {
// 		channel1 <- i
// 	}
// 	close(channel1) //after sending, closing channel1

// 	//getting result
// 	for i:=1; i<=noofjobs;i++{
// 		res := <- channel3
// 		fmt.Println(res)
// 	}

// }

package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {

    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= numJobs; a++ {
        <-results
    }
}