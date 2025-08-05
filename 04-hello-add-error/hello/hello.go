package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, err := greetings.Hello("Graldys")
	message, err := greetings.Hello("")

	// If an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
