package main

import (
	"fmt"
	"time"
)

type SignalState int

const (
	Red SignalState = iota
	Green
	Yellow
)

var SignalStateName = map[SignalState]string{
	Red:    "red",
	Green:  "green",
	Yellow: "yellow",
}

func (s SignalState) String() string {
	return SignalStateName[s]
}

func SignalTransition(s SignalState) SignalState {
	switch s {
	case Red:
		return Green
	case Green:
		return Yellow
	case Yellow:
		return Red
	default:
		panic(fmt.Errorf("unknown signal state: %d", s))
	}
}

func main() {
	current := Red

	for i := 0; i < 6; i++ {
		fmt.Println("Current Light:", current)
		time.Sleep(1 * time.Second) 
		current = SignalTransition(current)
	}
}
