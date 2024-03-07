package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returs a greeting given the person's name.
// Capitalized function name means it can be exported.
func Hello(name string) (string, error) {
	// If name is an empty string, return an error with message. 
	if name == "" {
		return "", errors.New("Name cannot be an empty string!")
	}

	// Return a greeting that embeds the name in the message.
	return fmt.Sprintf(randomFormat(), name), nil
}

// Return a map that associates each person to a greeting message.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	// Go through each name, generate a message for it.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// Return a random message format
func randomFormat() string {
	// Slice of formats
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a random format with random index
	return formats[rand.Intn(len(formats))]
}
