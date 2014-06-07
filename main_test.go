package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	result := Greeting()
	if result != "Hello world!" {
		t.Errorf("result is %s", result)
	}
}
