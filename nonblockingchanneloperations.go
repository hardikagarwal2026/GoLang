package main

import "fmt"

//Basically we can write non blocking send and receive with the help of seect plus default in it

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	//Non-Blocking Receive, as no routine sends to it , inspite of blocking, it runs default 
	// If a value is available on channel1 then select will take the <-channel1 case with that value. If not it will immediately take the default case.
	select {
	case msg:= <- channel1:
		fmt.Println("Message:", msg)
	default:
		fmt.Println("No message Received")
	}


	//Non-Blocking Send
	//msg cannot be sent to the channel1, because the channel1 has no buffer and there is no receiver.
	msg := "hi"
	select {
	case channel1 <- msg:
		fmt.Println("Message")
	default:
		fmt.Println("Hardik the great")
	}

	//Multiple Non-Blocking Receive, but no routine is sending is sending to channel1 and channel2, so default case executes
	select {
		case msg := <- channel1:
			fmt.Println("One", msg)
		case msg2 := <- channel2:
			fmt.Println("Two", msg2)
		default:
			fmt.Println("Janhvi")
	}

}