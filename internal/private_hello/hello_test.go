package private_hello

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello World!"
	actual := hello("World")
	if actual != expected {
		t.Errorf("HelloWorld() = %q, want %q", actual, expected)
	}
}
