package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, error := Hello("Gladys")
	if !want.MatchString(msg) || error != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, error, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if err == nil {
		t.Errorf("Hello(\"\") = %q, %v, 期待されるエラーが返されませんでした", msg, err)
	}
}
