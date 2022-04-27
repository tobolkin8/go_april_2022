package integers

import (
	"hello/tests"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	tests.FailedTestResults(t, sum, expected)
}
