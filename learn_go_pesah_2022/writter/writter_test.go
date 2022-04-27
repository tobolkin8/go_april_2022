package writter

import (
	"bytes"
	"fmt"
	"hello/str_consts"
	"hello/tests"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, str_consts.WRITTER_TEST_GREET_NAME)

	got := buffer.String()
	_,want := fmt.Printf(str_consts.WRITTER_GREET,str_consts.WRITTER_TEST_GREET_NAME)

	tests.FailedTestResults(t, got, want)
}
