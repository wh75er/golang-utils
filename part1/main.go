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

//    fmt.Println("first element:<", string(data[0][0]), ">")
//    fmt.Println("sec element:<", string(data[1][0]), ">")

//	fmt.Println("Result:\n", strings.Join(data, "\n"))

    e  = iop.WriteInput(data, iopts)

    if e != nil {
        log.Fatal(e)
    }
}
