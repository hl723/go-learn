package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    // log.SetFlags(0)

	// Get a greeting message from my module and print it.
	message, err := greetings.Hello("Hao")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	// Test Hellos, pass slice of names
	names := []string {"Hao1", "Hao2", "Hao3"}
	messages, err := greetings.Hellos(names)
	
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
	
	// Fail test
	message, err = greetings.Hello("")
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)
}
