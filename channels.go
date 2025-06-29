package main

import "fmt"

func main() {
 messages := make(chan string) /// creating the channel


 //sender goroutine
 go func(){ messages <- "Hardik the great"}() // we are passing the message to channel

//receiver goroutine
go func(){
    msg:= <- messages
	fmt.Println(msg)
}()

//select {} blocks the main goroutine forever, so the sender and receiver goroutines have time to run.
select {}

}