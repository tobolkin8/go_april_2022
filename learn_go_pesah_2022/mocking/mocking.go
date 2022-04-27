package mocking

import (
	"bytes"
	"fmt"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out *bytes.Buffer) {
	    for i := countdownStart; i > 0; i-- {
        time.Sleep(1 * time.Second)
        fmt.Fprintln(out, i)
    }

    time.Sleep(1 * time.Second)
    fmt.Fprint(out, finalWord)
}
