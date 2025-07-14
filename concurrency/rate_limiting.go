// //rate limiting is for efficient utilization and controlling of resources.
// // we can achieve this through the help of goroutines, channels and tickers in golang
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
//   //creating a buffered channel for 5 requests
//   requests := make(chan int, 5)
//   for i := 1; i<=5; i++ {
// 	requests <- i
//   }
//   close(requests) //closing channel, means no more request will be sent to the channel


//   //time.Tick creates a channel that ticks after every 200ms, so this is our rate limiter, one tick means one request allowed
//   limiter := time.Tick(200 * time.Millisecond)


//   for req := range requests {
// 	//By blocking on a receive from the limiter channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.
// 	<- limiter //to receive the tick, Wait until the limiter sends a tick (i.e., wait 200ms).
// 	fmt.Println("request", req, time.Now())
//   }

//   //creating a burstylimiter channel so that it can initially handle 3 requests, and then on e request at each tick
//    burstyLimiter := make(chan time.Time, 3)

//   //filling up the channel initially, to represent allowed bursting
//   	for i := 0; i < 3; i++ {
//     	burstyLimiter <- time.Now()
// 	}


//   //Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
//   go func() {
// 	for t := range time.Tick(200* time.Millisecond){
// 		burstyLimiter <- t
// 	}
//   }

//   //5 more incoming requests, first 3 three will burst
//   burstyRequests := make(chan int, 5)
//   for i:=1;i<=5;i++ {
// 		burstyRequests <- i
//   }
//   close(burstyRequests)

//   for req := range burstyRequests {
// 	<- burstyLimiter //Each <-burstyLimiter blocks until a new token becomes available.
// 	fmt.Println("request", req, time.Now())
//   }

// }

package main

import (
    "fmt"
    "time"
)

func main() {

    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    burstyLimiter := make(chan time.Time, 3)

    for range 3 {
        burstyLimiter <- time.Now()
    }

    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}