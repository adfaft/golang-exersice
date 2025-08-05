package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("hello")
	fmt.Println(quote.Go())
	name := "asdfasf"
	n := 98.98980908677567
	message := fmt.Sprintf("name %v %#v %T %10.5f %3.2f", name, name, name, n, n)
	fmt.Println(message)
}
