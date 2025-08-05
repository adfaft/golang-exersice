package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name
// and checking the valid return value
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Errorf("Hello('%v'}) = %q, %v, want match for %#q, nil", name, msg, err, want)
	}
}

func TestHelloNameEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf("Hello('') = %q, %v, want '', error", msg, err)
	}
}
