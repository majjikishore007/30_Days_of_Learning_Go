package main

import (
	"bytes"
	"testing"
)

func assertText(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got = %s, Want = %s", got, want)
	}
}

func TestLogger(t *testing.T) {

	t.Run("Test logger Info", func(t *testing.T) {
		buffer := bytes.Buffer{}
		logger := NewLogger(&buffer)
		logger.Info("info")
		got := buffer.String()
		want := "2025-07-03 16:47:40 [INFO] info"
		assertText(t, got, want)

	})
	t.Run("Test logger Debug", func(t *testing.T) {
		buffer := bytes.Buffer{}
		logger := NewLogger(&buffer)
		logger.Debug("info")
		got := buffer.String()
		want := "2025-07-03 16:47:40 [DEBUG] info"
		assertText(t, got, want)

	})
	t.Run("Test logger Error", func(t *testing.T) {
		buffer := bytes.Buffer{}
		logger := NewLogger(&buffer)
		logger.Error("info")
		got := buffer.String()
		want := "2025-07-03 16:47:40 [ERROR] info"
		assertText(t, got, want)

	})
}
