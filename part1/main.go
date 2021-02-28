package main

import (
	"fmt"
	"log"
	"uniq/args"
	"uniq/iop"
)

const usage = "Usage: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]"

func main() {
	opts, iopts, e := args.ParseArgs()

	if e != nil {
		log.Fatal(e, "\n", usage)
	}

	e = opts.IsValid()

	if e != nil {
		log.Fatal(e, "\n", usage)
	}

	data, e := iop.ReadInput(iopts)

	if e != nil {
		log.Fatal(e)
	}
}
