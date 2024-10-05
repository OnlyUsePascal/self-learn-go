package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Joun"
	// msgExpected := regexp.MustCompile("\b" + name + "\b") // somehow using this will fail the test.
	msgExpected := regexp.MustCompile(`\b` + name + `\b`)
	msgActual, err := Hello(name)

	// epxected matching result + no error
	if !msgExpected.MatchString(msgActual) || err != nil {
		t.Fatalf("Hello(%q) returns %q, %v. Expect a match for %#q, nil", name, msgActual, err, msgExpected)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	msgExpected := ""
	msgActual, err := Hello(name)

	// expect blank return + error
	if msgActual != msgExpected || err == nil {
		t.Fatalf(`Hello(%q) returns %q, %v. Expect "", error`, name, msgActual, err)
	}
}
