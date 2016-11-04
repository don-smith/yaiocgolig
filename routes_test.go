package main

import "testing"

func TestHello(t *testing.T) {
	var expected = "<h1>Hello, world!</h1>"
	if Hello() != expected {
		t.Error("unexpected output")
	}
}
