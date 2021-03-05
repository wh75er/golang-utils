package main

import (
	"fmt"
	"os"
	"uniq/args"
	"uniq/iop"
	"uniq/unique"
)

const usage = "Usage: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]"

func main() {
	opts, iopts, e := args.ParseArgs()
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	e = opts.IsValid()
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	data, e := iop.ReadInput(iopts)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	data, e = unique.Uniqualize(data, opts)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	e = iop.WriteInput(data, iopts)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
