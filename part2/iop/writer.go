package iop

import (
	"bufio"
	"fmt"
	"io"
)

func WriteInput(s string, out io.Writer) {
	w := bufio.NewWriter(out)
	fmt.Fprintln(w, s)
	w.Flush()
}
