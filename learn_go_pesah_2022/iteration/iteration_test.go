package iteration

import (
	"hello/tests"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	for i := 0; i < 5; i++ {
		repeated := Repeat("a", i)
		expected := strings.Repeat("a", i)

				tests.FailedTestResults(t, expected, repeated)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", i)
	}
}
