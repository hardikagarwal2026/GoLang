package main

import (
	"fmt"
)

func ping(msg string, channel1 chan<- string){
     channel1 <- msg
}

func pong(channel1 <-chan string, channel2 chan<- string ){
	message := <- channel1
	channel2 <- message
}

func main(){
	pings := make(chan string)
	pongs := make(chan string)

	go ping("hardik", pings)
	go pong(pings, pongs)

	fmt.Println(<-pongs)
}