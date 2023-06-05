package helpers

import (
	"os"
	"testing"
)

func TestGetEnvOrUseDefaultValue(t *testing.T) {
	t.Run("Should load env with correct value", func(t *testing.T) {
		os.Setenv("TEST_ENV_VAR", "test value")

		got := GetEnvOrUseDefaultValue("TEST_ENV_VAR", "default value")
		want := "test value"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		os.Unsetenv("TEST_ENV_VAR")
	})

	t.Run("Should use default value on non exist env", func(t *testing.T) {
		got := GetEnvOrUseDefaultValue("NOT_TEST_ENV_VAR", "default value")
		want := "default value"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
