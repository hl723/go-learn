package greetings

import (
	"testing"
	"regexp"
)

// TestHelloName calls greeting.Hello with a name, normal case
func TestHelloName(t *testing.T) {
	name := "Hao"
	
	// MustCompile will fail if bad pattern
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Hao") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greeting.Hello with empty name, checking for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
