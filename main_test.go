package main

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	os.Setenv("TEST_ENV_VAR", "test value")

	LoadEnv()

	got, exists := os.LookupEnv("TEST_ENV_VAR")
	want := "test value"

	if !exists {
		t.Error("expected TEST_ENV_VAR toe exist")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	os.Unsetenv("TEST_ENV_VAR")
}

func TestNewServer(t *testing.T) {
	os.Setenv("PORT", "8000")

	server := NewServer()

	got := server.Port
	want := "8000"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	os.Unsetenv("PORT")
}
