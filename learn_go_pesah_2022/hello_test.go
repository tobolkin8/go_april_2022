package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hi ya`ll"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
