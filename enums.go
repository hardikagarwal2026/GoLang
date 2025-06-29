package main

import "fmt"

//defining a enum
type ServerState int

const(
	StateIdle ServerState = iota
    StateConnected
	StateError
	StateRetrying
)

var statename = map[ServerState] string {
	StateIdle : "idle",
    StateConnected: "connected",
	StateError: "error",
	StateRetrying: "retrying",
}

func (ss ServerState) String() string {
    return statename[ss]
}


func transition(ss ServerState) ServerState {
	switch ss{
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	default:
		panic(fmt.Errorf("Unknown Errror"))
	}
}

func main() {
    ns := transition(StateIdle)
	fmt.Println(ns)

	ns2:= transition(ns)
	fmt.Println(ns2)
}
