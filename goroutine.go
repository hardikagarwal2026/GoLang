package main

import (
	"fmt"
	"time"
)

func f(msg string) {
	for i:= range 3{
		fmt.Println(msg, ":", i)
	}
}

func main(){
	f("Janhvi Agarwal")

	//running function asynchronous in go rourtine,
	go f("Lord Hardik")


	//running asynchronous function in go routine, output mayb be interleaved
	go func(message string){
		fmt.Println(message)
	}("Lord Hardik the great")

	time.Sleep(time.Second)
	fmt.Println("done")


}