package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// messages, err := greetings.Hellos([]string{"fauzi", "ezi"})

	names := []string{"fauzi", "ezi"}
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
