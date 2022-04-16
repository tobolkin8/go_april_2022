package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("hello-success", func(t *testing.T) {
		for language := range Languages {
			// what we are getting from the real func
			got := Hello("Nina", language)
			// what we expect to get from the real func
			want := ReturnLangText(language) + "Nina"
			if got != want {
				assertCorrectMessage(t, got, want)
			}
		}

	})

	t.Run("hello-fail", func(t *testing.T) {
		// what we are getting from the real func
		got := Hello("", "")
		// what we expect to get from the real func
		want := "Hello, world!"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
}
