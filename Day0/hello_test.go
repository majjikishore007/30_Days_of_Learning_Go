package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := say_hello("Kishore", "English")
		want := "Hello, Kishore"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello world default case", func(t *testing.T) {
		got := say_hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in spanish", func(t *testing.T) {
		got := say_hello("Kishore", "Spanish")
		want := "Hola, Kishore"
		assertCorrectMessage(t, got, want)
	})
}
