package main

import (
	"fmt"
	"time"
)

//select helps to wait on multiple goroutines

func main() {
	ch1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2) //it waits for 2 sec, before sending message to channel1
		ch1 <- "hardik"
	}()
	select{ //ch1 takes 2 sec, while another case takes 1 sec, so it will print second case
	case res:= <- ch1:
		fmt.Println(res)
	case <- time.After(1 * time.Second):
		fmt.Println("TimeOut")
	}


	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- "Hogaya"
	}()
	select { //in this select runs that case which has less timeout	
	case res:= <- ch2:
		fmt.Println(res)
	case <- time.After(time.Second * 2):
		fmt.Println("The Great")
	}
}



