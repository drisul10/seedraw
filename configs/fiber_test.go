package configs

import (
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestNewFiberConfig(t *testing.T) {
	got := NewFiberConfig()
	want := fiber.Config{}

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
