package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello return a simple randomized welcome message
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	format := randomMessageFormat()
	message := fmt.Sprintf(format, name)
	return message, nil
}

// Hellos return multiple randomized welcome messages for multiple names
func Hellos(names []string) (map[string]string, error) {

	// A map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, error := Hello(name)

		if error != nil {
			return nil, error
		}

		messages[name] = message
	}

	return messages, nil
}

func randomMessageFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
