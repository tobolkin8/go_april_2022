package writter

import (
	"bytes"
	"fmt"
	"hello/str_consts"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, str_consts.WRITTER_GREET, name)
}
