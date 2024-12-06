package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// give error if name is blank
	if name == "" {
		return "", errors.New("At least state your name!")
	}

	// var message string
	// message = fmt.Sprintf("G' Morning, %v!",name)
	format := randomFormat()
	message := fmt.Sprintf(format, name)
	return message, nil
}

func randomFormat() string {
	// format slices
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a random selected format
	return formats[rand.Intn(len(formats))]
}
