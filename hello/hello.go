package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	// set properties of predefined Logger, including the log entry prefix and a flag to disable printing the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"Gladdie", "SAM", "Fuji-san"}

	// request a greeting message.
	message, err := greetings.Hellos(names)
	// if error was returned, print it to console and exit program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message to console
	fmt.Println(message)

	fmt.Println("\nHello from go.")
	fmt.Println(quote.Glass())
}
