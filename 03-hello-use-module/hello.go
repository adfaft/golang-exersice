package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Graldys")
	fmt.Println(message)
}
