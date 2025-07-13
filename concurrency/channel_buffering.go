package main


import "fmt"


func main(){
messages:= make(chan string, 3) //creating a buffered channel with 3 values, in buffered channel, it should have receiver
messages <- "hardik"
messages <- "the"
messages <- "great"
fmt.Println(<- messages)
fmt.Println(<- messages)
fmt.Println(<- messages)

}