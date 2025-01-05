package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greetings for the named person

func Hello(name string) (string, error) {
	// Return a greeting thst embeds the name in a message.
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Craete a message using a random format.

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

// Hellos returns a map that associates each of the named people
// with a greeting message

func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages.

	messages := make(map[string]string)

	// loops through the received slices of names, calling
	// the Hello function to get a message for each name

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		//  =in the map, associate the retrieved message with
		// the name

		messages[name] = message
	}
	return messages, nil
}

// randomFormat retruns one of a set of greeting messages. The returned
// message is selected at random

func randomFormat() string {
	// A slice of message formats.

	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v!",
		"Hail, %v! Well met",
	}

	// return a randomly selected message format by specifying
	// a random index for the slice of the formats.

	return formats[rand.Intn(len(formats))]
}
