package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Joun")
	fmt.Println(message)
}
