package dependency_injection

import (
	"testing"
	"bytes"
	"os"
)


func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	Greet(os.Stdout , "Elodie")
}