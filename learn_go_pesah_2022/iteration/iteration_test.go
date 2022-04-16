package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	for i := 0; i < 5; i++ {
		repeated := Repeat("a", i)
		expected := strings.Repeat("a", i)

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", i)
	}
}
