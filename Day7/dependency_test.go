package main

import (
	"bytes"
	"testing"
)

// buffer is an array of bytes that impltes Writer and Reader interfaces
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kishore")
	got := buffer.String()
	want := "Hello, Kishore"

	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}
