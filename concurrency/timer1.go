//In this program, we simulate a situation where we’re waiting for an operation (like a network call) to complete, but we don’t want to wait more than 5 seconds. We start a 5-second timer and also launch a goroutine that simulates the operation by sleeping for 2 seconds. Using a select, we wait for either the operation to finish (done channel) or the timer to fire (timer.C). If the operation finishes first, we stop the timer to avoid it firing later; if it already fired (which is rare here), we drain its channel to avoid blocking. This is a common pattern in Go to implement timeouts with early cancellation when operations finish quickly.
package main

///Let’s say you’re waiting for a response from a network call or some operation, but you want to timeout after 5 seconds if it takes too long. However, if the operation finishes in 2 seconds, you don’t need the timer anymore

import (
    "fmt"
    "time"
)

func main() {
	
    timer := time.NewTimer(5 * time.Second)

    done := make(chan bool)

    go func() {
        // Simulate fast operation
        time.Sleep(2 * time.Second)
        done <- true
    }()

    select { //select checks for which routineis available first
    case <-done:
		//The timer has not fired yet — meaning the timer is still running and its channel timer.C has not received any value, if it reutrns true value
        if timer.Stop() {
            fmt.Println("Finished early, timer stopped")
        } else {
            // Already fired, maybe need to drain it
            <-timer.C
        }
    case <-timer.C:
        fmt.Println("Timeout occurred")
    }
}

//If the timer hasn't fired yet (it hasn't — it's only been 2 seconds), Stop() returns true, meaning we successfully cancelled the timer.