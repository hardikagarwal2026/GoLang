package main

import (
	"fmt"
	"time"
)

//created a worker go routine takes channel as a parameter and sends, a response as true after done working
func worker(done chan bool){
	fmt.Println("Starting...")
	time.Sleep(time.Second)
	fmt.Println("Done...")
	done <- true
}


func main() {
	hardik := make(chan bool, 1) //created a channel,
	go worker(hardik) // passed channel as a parameter

	//blocks until we receive the notification from the receiver, so it is called blocking receive
	<- hardik // gets a result in the channel, also as we have created a buffer channel if we remove the <- hardik , then program exits even before the worker even completed

}