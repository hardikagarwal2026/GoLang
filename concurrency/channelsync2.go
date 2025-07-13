package main

import (
	"fmt"
)

//square function taking int and channel as parameter passinf resuklt to the channel
func square(i int, channel chan int){
	result := i*i
	channel <- result
}

func main() {
   a := []int{1,2,3}
   hardik := make(chan int) //creating channel

//worker sends back its result via a channel.
   for _, n := range a {
		go square(n, hardik) //passing int and channel
   }

   for range a{
	   fmt.Println(<- hardik) //<-hardik it recives the output from the gorouitine, blocks the main unitl it receives for all the goroutines
   }
}