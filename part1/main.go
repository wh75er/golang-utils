package main

import (
	//	"fmt"
	"log"
	//	"strings"
	"uniq/args"
	"uniq/iop"
	"uniq/unique"
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

	data, e = unique.Uniqualize(data, opts)

	if e != nil {
		log.Fatal(e)
	}

//	fmt.Println("last element:<", data[len(data) - 1], ">")

    e  = iop.WriteInput(data, iopts)

    if e != nil {
        log.Fatal(e)
    }
}
