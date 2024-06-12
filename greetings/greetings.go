package greetings // declare a greetings package to collect related functions

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) { // implemented a Hello() function to return a greeting
	// if no name was given, return an error with message
	if name == "" {
		return "", errors.New("empty name")
	}

	// if name was received, return a value that embeds the name in a message
	//message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprint(randomFormat())
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)
	// loop through the received slice of names, calling the Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a rand selected message format by specifying a rand index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
