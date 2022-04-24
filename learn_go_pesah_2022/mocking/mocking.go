package mocking

import (
	"bytes"
	"fmt"
)

func Countdown(out *bytes.Buffer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")
}
