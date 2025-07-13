package main

import (
	"fmt"
	"time"
)

func worker(done chan bool){
	fmt.Println("Starting...")
	time.Sleep(time.Second)
	fmt.Println("Done...")
	done <- true
}


func main(){
    hardik := make(chan bool)

	//starting 3 workers concurrently
	for i:=0;i<3;i++ {
		go worker(hardik) //passing hardik channel as a parameter
	}

	for i:=0;i<3;i++ {
		<- hardik //blocking receive , it waits for all the workers to finish before exit
	}

}