package greetings_test

import (
	"regexp"
	"testing"

	"example.com/greetings"
)

func TestHelloName(t *testing.T) {
	name := "Joun"
	target := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := greetings.Hello(name)
	
	if !target.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(%s) outputs [%q, %v], expecting [%#q, nil]`, name, msg, err, target)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	msg, err := greetings.Hello(name)

	if msg != "" || err == nil {
		t.Fatalf(`Hello(%s) outputs [%q, %v], expecting ["", error]`, name, msg, err)
	}
}