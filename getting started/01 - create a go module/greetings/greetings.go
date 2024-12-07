package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Bro missing name :(")
	}

	message := fmt.Sprintf(randomFormat(), name)

	// or this
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome Home!", name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	
	for _, name := range names {
		if name == "" {
			return nil, errors.New("Bro missing one of the names :(")
		}

		message := fmt.Sprintf(randomFormat(), name)
		messages[name] = message
	}
	
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome Home!!",
		"Great to see you, %v!",
	}

	return formats[rand.Intn(len(formats))]
}
