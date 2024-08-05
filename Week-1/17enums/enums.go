package main

import "fmt"

// Define a new type named ServerState with underlying int type
type ServerState int

// Define constants for the ServerState type using iota to generate successive values
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Map to associate ServerState values with their string representations
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// Implement the String method for ServerState to convert enum values to strings
func (ss ServerState) String() string {
	return stateName[ss]
}

// transition function emulates a state transition for a server
// It takes the current state and returns the next state
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Here you can add conditions to determine the next state
		return StateIdle
	case StateError:
		return StateError
	default:
		// Panic if an unknown state is encountered
		panic(fmt.Errorf("unknown state: %s", s))
	}
	// Default return (though it will never be reached)
	return StateConnected
}

func main() {
	fmt.Println("---Enums---")
	// Start with StateIdle and transition to the next state
	ns := transition(StateIdle)
	fmt.Println(ns)

	// Transition from the new state
	ns2 := transition(ns)
	fmt.Println(ns2)
}
