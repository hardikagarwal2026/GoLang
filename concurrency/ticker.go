// A time.Ticker is used when you want to do something repeatedly at regular intervals

package main

import (
    "fmt"
    "time"
)

func main() {
    //time.NewTicker(1 * time.Second) creates a ticker that sends the current time every 1 second on the channel ticker.C
	ticker := time.NewTicker(1 * time.Second)

    done := make(chan bool)

    go func() {
        time.Sleep(5 * time.Second)
        done <- true
    }()

    for {
        select {
        case <-done:
            ticker.Stop()
            fmt.Println("Ticker stopped")
            return
        case t := <-ticker.C:
            fmt.Println("Tick at", t)
        }
    }
}
