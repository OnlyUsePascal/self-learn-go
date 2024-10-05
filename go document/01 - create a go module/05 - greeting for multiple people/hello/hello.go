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

	// name := "Joun"
	// message, err := greetings.Hello(name)
	// exit if error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// otherwise print casually
	// fmt.Println(message)

	names := []string{"Joun", "Jeff", "Jay"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
