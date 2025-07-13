// program to demostrate multiple goroutine passing and processing data through multiple channels

// package main

// import "fmt"

// func A(channel1 <-chan int,channel2 chan<- int ){
// 	val := <- channel1
//  	channel2 <- val *  2
// }

// func B(channel2 <-chan int, channel3 chan<- int){
// 	val := <- channel2
//  	channel3 <- val *  2
// }


// func main(){
//    channel1 := make(chan int)
//    channel2 := make(chan int)
//    channel3 := make(chan int)


//    go A(channel1, channel2)
//    go B(channel2, channel3)
   

//    channel1 <- 1

//    result := <- channel3
//    fmt.Println(result)

// }


package main

import "fmt"

func A(channel1 <-chan int,channel2 chan<- int ){
	for val := range channel1 { //receiving data from channel1
		channel2 <- val *2
	}
	close(channel2) //close after all data sent
}

func B(channel2 <-chan int, channel3 chan<- int){
	for val := range channel2{ //receiving data from channel2
		channel3 <- val * 2
	}
	close(channel3) //close after all data sent
}


func main(){
	a := []int{1,2,3,4,5}

   channel1 := make(chan int, len(a)) //creating  buffered channels with capacity length of a
   channel2 := make(chan int, len(a))
   channel3 := make(chan int, len(a))


   go A(channel1, channel2)
   go B(channel2, channel3)
   
   for _, num := range a {
       channel1 <- num
   }
   close(channel1) //close after all data sent

   for range a {
	fmt.Println("Result:", <- channel3) //final receive through channel3
   }

}