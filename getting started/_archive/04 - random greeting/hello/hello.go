package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// configure logger
	log.SetPrefix("greetings: ")
	log.SetFlags(1)

	name := "Joun"
	message, err := greetings.Hello(name)
	// exit if error
	if err != nil {
		log.Fatal(err)
	}

	// otherwise print casually
	fmt.Println(message)
}
