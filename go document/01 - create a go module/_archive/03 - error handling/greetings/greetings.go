package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// give error if name is blank
	if name == "" {
		return "", errors.New("At least state your name!")
	}

	// var message string
	// message = fmt.Sprintf("G' Morning, %v!",name)

	message := fmt.Sprintf("G' Morning, %v!", name)
	return message, nil
}
