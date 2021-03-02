package iop

import (
    "fmt"
	"bufio"
	"os"
	"uniq/args"
)

func WriteInput(data []string, opts args.IoOptions) error {
	outStream := os.Stdout

	if opts.OutFilename != "" {
		f, e := os.OpenFile(opts.OutFilename, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, os.FileMode(0644))

		if e != nil {
			return e
		}

		defer f.Close()


		outStream = f
	}

	w := bufio.NewWriter(outStream)

	for _, v := range data {
        fmt.Fprintln(w, v)
    }
    w.Flush()

	return nil
}
