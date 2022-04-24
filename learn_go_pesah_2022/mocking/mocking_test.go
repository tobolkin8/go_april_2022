package mocking

import (
	"bytes"
	"hello/tests"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	tests.FailedTestResults(t, got, want)
	tests.CorrectResults(got, want)
}
