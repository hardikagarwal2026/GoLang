package main

import "fmt"
import "time"

func main(){
	channel1 := make(chan string)
	channel2 := make(chan string)
	go func(){
		time.Sleep(1 * time.Second)
		channel1 <- "hardik"
	}()

	go func ()  {
		time.Sleep(1 * time.Second)
		channel2 <- "the great"
	}()

	for range 2{
		// select lets us wait on multiple channel operations
		select {
		case msg1 := <- channel1:
			fmt.Println("Received:", msg1)
		case msg2 := <- channel2:
			fmt.Println("Received:", msg2)
		}
	}
}