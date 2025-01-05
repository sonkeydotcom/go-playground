package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// Set properties of the predefined Logger, incuding
	// the log entry prefix and a flag to diable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names.
	names := []string{"Gladys", "Samantha", "Darin"}

	// Request a greeting message

	messages, err := greetings.Hellos(names)
	// if an error was returned, print it into the console and
	// exit the application

	if err != nil {
		log.Fatal(err)
	}

	// If no errro was returned, print the returned message
	// to the console

	fmt.Println(messages)
}
