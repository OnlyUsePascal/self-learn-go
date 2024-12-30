package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
	}

	for _, tc := range testcases {
		got, _ := Reverse(tc.in)

		if got != tc.want {
			t.Errorf("Reverse(%q) = %q; want %q", tc.in, got, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}

	// you don't have control over input, hence the expected output
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, in string) {
		rev, revErr := Reverse(in)
		if revErr != nil {
			t.Skip()
		}

		doubleRev, doubleRevErr := Reverse(rev)
		if doubleRevErr != nil {
			t.Skip()
		}	

		if in != doubleRev {
			t.Errorf("Before: %q, after: %q", rev, doubleRev)
		}
		if utf8.ValidString(in) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
