package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	fmt.Println("=== One name ===")
	name := "Joun"
	message, err := greetings.Hello(name)

	if err != nil {
		// using Fatal will print the error and exit program
		log.Fatalln(err)
	}

	fmt.Println(message)
	

	fmt.Println("=== Many names ===")	
	names := []string{"joun", "james", "sanctuary"}
	messages, err := greetings.Hellos(names)
	
	if err != nil {
		// using Fatal will print the error and exit program
		log.Fatalln(err)
	}
	
	fmt.Println(messages)
}
